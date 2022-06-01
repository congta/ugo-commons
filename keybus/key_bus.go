package keybus

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/congta/ugo-commons/commons-codec/ucodings"
	"github.com/congta/ugo-commons/commons-codec/udigests"
	"github.com/congta/ugo-commons/commons-codec/usecrets"
	"github.com/congta/ugo-commons/commons-lang/ubytes"
)

/**
TODO 拆到独立的项目里
version = 0：这个版本没有 version 字段，是原 Java 代码生成的串
version = 1：go 的第一版，aes cbc 加解密：version(1 byte) + holderId(1 byte) + headerLen(2 byte) + 密文，
最大可使用 127 个 holder，holderId 负数保留，可用作扩展场景，headerLen 是原文中 json header 的长度
*/

var (
	version     byte = 1
	minHolderId      = 0
	maxHolderId      = 127
)

type KeyBus interface {
	Encrypt(data []byte) (res []byte, err error)
	EncryptStr(data string) (res string, err error)
	Decrypt(data []byte) (res []byte, err error)
	DecryptStr(data string) (res string, err error)
}

type KeyHolder struct {
	Id  int
	Key []byte
	Iv  []byte
}

func (h KeyHolder) ToString() string {
	return fmt.Sprintf("%s~%s~%d", ucodings.EncodeBase64String(h.Key), ucodings.EncodeBase64String(h.Iv), h.Id)
}

type Algorithm int8

var (
	AesCBC Algorithm = 0
	AesCTR Algorithm = 1
)

type MsgHeader struct {
	Time int64      `json:"t"`
	Algo *Algorithm `json:"a,omitempty"`
}

func encrypt(data []byte, holder KeyHolder) (res []byte, err error) {
	if holder.Id < minHolderId || holder.Id > maxHolderId {
		return nil, errors.New("invalid holder id for current version")
	}

	// add outer header (unencrypted)
	finalBytes := make([]byte, 0)
	finalBytes = append(finalBytes, version)
	finalBytes = append(finalBytes, byte(holder.Id))

	// add inner header (encrypted)
	header := MsgHeader{
		Time: time.Now().UnixMilli(),
		// Algo: &AesCBC,
	}
	headerBytes, _ := json.Marshal(header)
	headerLen := len(headerBytes)

	iv := getIv(holder, headerBytes)
	msgBytes, err := usecrets.AesCBCEncrypt(data, holder.Key, iv)
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
	header := MsgHeader{}
	if err = json.Unmarshal(data[4:msgStart], &header); err != nil {
		return nil, err
	}

	holder, ok := holderMap[id]
	if !ok {
		return nil, fmt.Errorf("holder not found: %d", id)
	}

	iv := getIv(holder, data[4:msgStart])
	if header.Algo == nil {
		header.Algo = &AesCBC
	}
	algo := *header.Algo
	switch algo {
	case AesCBC:
		dst, err = usecrets.AesCBCDecrypt(data[msgStart:], holder.Key, iv)
	case AesCTR:
		err = fmt.Errorf("algorithm AES CTR not supported")
	default:
		err = fmt.Errorf("unknown algorithm %d", algo)
	}
	return dst, err
}

func getIv(holder KeyHolder, header []byte) []byte {
	iv := holder.Iv
	if iv == nil || len(iv) == 0 {
		iv = udigests.Md5(header)
	}
	return iv
}
