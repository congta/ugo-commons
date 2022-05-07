package usecrets

import (
	"congta.com/ugo-commons/commons-codec/ucodings"
	"congta.com/ugo-commons/commons-io/ufiles"
	"congta.com/ugo-commons/commons-lang/unumbers"
	"congta.com/ugo-commons/commons-logging/ulogs"
	"congta.com/ugo-commons/commons-u/ucommons"
	"fmt"
	"strings"
)

type KeyBoxByFile struct {
	holderMap map[int]KeyHolder
	holderArr []KeyHolder
	arrCursor int
}

func NewKeyBoxByFile(sid string) *KeyBoxByFile {
	fileName := GetKeyBoxFileName(sid)
	lines, err := ufiles.ReadLines(fileName)
	if err != nil {
		ulogs.Panic("key center secret not ready for %s", sid)
	}

	holders := make(map[int]KeyHolder)
	holderArray := make([]KeyHolder, 0, len(lines))
	for _, keyStr := range lines {
		if !strings.Contains(keyStr, `~`) {
			continue
		}
		ki := strings.Split(keyStr, "~")
		holder := &KeyHolder{}
		if len(ki) > 2 {
			holder.Id = unumbers.ParseIntWildly(ki[2])
		}
		holder.Key = ucodings.DecodeBase64StringWildly(ki[0])
		holder.Iv = ucodings.DecodeBase64StringWildly(ki[1])
		if _, ok := holders[holder.Id]; ok {
			ulogs.Warn("duplicate secret id: %d, use the first one", holder.Id)
			continue
		}
		holders[holder.Id] = *holder
		holderArray = append(holderArray, *holder)
	}
	box := &KeyBoxByFile{
		holderMap: holders,
		holderArr: holderArray,
	}
	if len(holderArray) < 1 {
		ulogs.Panic("no key found for KeyBox %s", fileName)
	}
	return box
}

func GetKeyBoxFileName(sid string) string {
	var fileName string
	if ucommons.IsWindows() {
		fileName = fmt.Sprintf(ucommons.Home()+"/.congta/key/%s.keb", sid)
	} else {
		fileName = fmt.Sprintf("/etc/conf/congta/key/%s.keb", sid)
	}
	return fileName
}

func (t *KeyBoxByFile) Encrypt(data []byte) (res []byte, err error) {
	holder := t.holderArr[t.arrCursor]
	t.arrCursor = (t.arrCursor + 1) % len(t.holderArr)

	return encrypt(data, holder)

}
func (t *KeyBoxByFile) EncryptStr(data string) (res string, err error) {
	secBytes, err := t.Encrypt([]byte(data))
	if err != nil {
		return "", err
	}
	return ucodings.EncodeBase64URLSafeString(secBytes), nil
}

func (t *KeyBoxByFile) Decrypt(data []byte) (res []byte, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = fmt.Errorf("%v", err0)
		}
	}()
	return decrypt(data, t.holderMap)
}

func (t *KeyBoxByFile) DecryptStr(data string) (res string, err error) {
	secBytes, err := ucodings.DecodeBase64String(data)
	if err != nil {
		return "", err
	}
	rawBytes, err := t.Decrypt(secBytes)
	if err != nil {
		return "", err
	}
	return string(rawBytes), nil
}
