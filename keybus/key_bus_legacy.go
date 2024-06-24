package keybus

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/congta/ugo-commons/commons-codec/usecrets"
	"github.com/congta/ugo-commons/commons-lang/ubytes"
)

// Java use legacy version

func encryptLegacy(data []byte, holder KeyHolder) (res []byte, err error) {
	if holder.Id < minHolderId || holder.Id > maxHolderId {
		return nil, errors.New("invalid holder id for current version")
	}

	finalBytes := make([]byte, 0)
	finalBytes = append(finalBytes, ubytes.IntToBytes(holder.Id)...)
	finalBytes = append(finalBytes, '~')

	if holder.Id > 0 {
		// add inner header (encrypted)
		header := MsgHeader{
			Time: time.Now().Unix(),
		}
		headerBytes, _ := json.Marshal(header)
		headerBytes = append(headerBytes, '~')
		data = append(headerBytes, data...)
	}

	var msgBytes []byte
	msgBytes, err = usecrets.AesCTREncrypt(data, holder.Key, holder.Iv)

	finalBytes = append(finalBytes, version)
	finalBytes = append(finalBytes, byte(holder.Id))
	finalBytes = append(finalBytes, msgBytes...)
	return finalBytes, nil
}

func decryptLegacy(data []byte, holderMap map[int]KeyHolder) (dst []byte, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = fmt.Errorf("%v", err0)
		}
	}()
	if data[4] != '~' {
		return nil, fmt.Errorf("failed to find the first separator")
	}

	id := int(binary.BigEndian.Uint32(data[:4]))
	if id < minHolderId || id > maxHolderId {
		return nil, fmt.Errorf("invalid holder id for current version: %d", id)
	}

	holder, ok := holderMap[id]
	if !ok {
		return nil, fmt.Errorf("holder not found: %d", id)
	}
	dst, err = usecrets.AesCTRDecrypt(data[5:], holder.Key, holder.Iv)

	if id > 0 {
		idx := bytes.IndexByte(dst, '~')
		//header := dst[:idx]
		dst = dst[idx+1:]
	}

	return dst, err
}
