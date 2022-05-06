package usecrets

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AesCTREncrypt(data, key, iv []byte) ([]byte, error) {
	// 1. create cipher.Block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 2. create stream
	stream := cipher.NewCTR(block, iv)
	// 3. encrypt
	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst, nil
}

func AesCTRDecrypt(data, key, iv []byte) ([]byte, error) {
	return AesCTREncrypt(data, key, iv)
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, pad...)
}

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// AesCBCEncrypt aes encryptor，len of key corresponding to algo, 16 is AES-128, 24 is AES-192, 32 is AES-256.
func AesCBCEncrypt(plainData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if blockSize != len(iv) {
		return nil, fmt.Errorf("iv len (%d) != block size (%d)", len(iv), blockSize)
	}
	plainData = PKCS7Padding(plainData, blockSize)
	cipherText := make([]byte, len(plainData))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainData)

	return cipherText, nil
}

func AesCBCDecrypt(cipherData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	plainData := make([]byte, len(cipherData))
	mode.CryptBlocks(plainData, cipherData)
	plainData = PKCS7UnPadding(plainData)
	return plainData, nil
}
