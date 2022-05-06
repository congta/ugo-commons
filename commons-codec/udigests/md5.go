package udigests

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src []byte) []byte {
	dst := md5.Sum(src)
	return dst[:]
}

func Md5String(src []byte) string {
	dst := Md5(src)
	return hex.EncodeToString(dst)
}
