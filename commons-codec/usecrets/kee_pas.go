package usecrets

import (
	"congta.com/ugo-commons/commons-codec/ucodings"
	"congta.com/ugo-commons/commons-codec/udigests"
	"congta.com/ugo-commons/commons-lang/ubytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

/**
version = 0：这个版本没有 version 字段，是原 Java 代码生成的串
version = 1：go 的第一版，aes cbc 加解密：version(1 byte) + holderId(1 byte) + headerLen(2 byte) + 密文，
最大可使用 127 个 holder，holderId 负数保留，可用作扩展场景，headerLen 是原文中 json header 的长度
*/

var (
	version     byte = 1
	minHolderId      = 0
	maxHolderId      = 127
)

type KeePas interface {
	encrypt(data []byte) (res []byte, err error)
	encryptStr(data string) (res string, err error)
	decrypt(data []byte) (res []byte, err error)
	decryptStr(data string) (res string, err error)
}

type KeyHolder struct {
	id  int
	key []byte
	iv  []byte
}

func (h KeyHolder) toString() string {
	return fmt.Sprintf("%s~%s~%d", ucodings.EncodeBase64String(h.key), ucodings.EncodeBase64String(h.iv), h.id)
}

func encrypt(data []byte, holder KeyHolder) (res []byte, err error) {
	if holder.id < minHolderId || holder.id > maxHolderId {
		return nil, errors.New("invalid holder id for current version")
	}

	// add outer header (unencrypted)
	finalBytes := make([]byte, 0)
	finalBytes = append(finalBytes, version)
	finalBytes = append(finalBytes, byte(holder.id))

	// add inner header (encrypted)
	header := make(map[string]interface{})
	header["t"] = time.Now().Second()
	headerBytes, _ := json.Marshal(header)
	headerLen := len(headerBytes)

	iv := getIv(holder, headerBytes)
	msgBytes, err := AesCBCEncrypt(data, holder.key, iv)
	if err != nil {
		return nil, err
	}

	finalBytes = append(finalBytes, ubytes.ShortToBytes(uint16(headerLen))...)
	finalBytes = append(finalBytes, headerBytes...)
	finalBytes = append(finalBytes, msgBytes...)
	return finalBytes, nil
}

func decrypt(data []byte, holderMap map[int]KeyHolder) (dst []byte, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = fmt.Errorf("%v", err0)
		}
	}()

	ver := data[0]
	if ver > version {
		return nil, fmt.Errorf("invalid holder header version %d", ver)
	}
	id := int(data[1])
	if id < minHolderId || id > maxHolderId {
		return nil, fmt.Errorf("invalid holder id for current version: %d", id)
	}

	headerLen := ubytes.BytesToInt16(data[2:4])
	msgStart := 4 + headerLen
	header := make(map[string]interface{})
	if err = json.Unmarshal(data[4:msgStart], &header); err != nil {
		return nil, err
	}

	holder, ok := holderMap[id]
	if !ok {
		return nil, fmt.Errorf("holder not found: %d", id)
	}

	iv := getIv(holder, data[4:msgStart])
	if dst, err = AesCTRDecrypt(data[msgStart:], holder.key, iv); err != nil {
		return nil, err
	}
	return dst, nil
}

func getIv(holder KeyHolder, header []byte) []byte {
	iv := holder.iv
	if iv == nil || len(iv) == 0 {
		iv = udigests.Md5(header)
	}
	return iv
}
