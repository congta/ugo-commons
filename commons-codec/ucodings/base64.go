package ucodings

import (
	"encoding/base64"
	"strings"

	"github.com/congta/ugo-commons/commons-logging/ulogs"
)

var (
	stdEncoding = base64.StdEncoding
	urlEncoding = base64.URLEncoding
)

func EncodeBase64(src []byte) []byte {
	dst := make([]byte, stdEncoding.EncodedLen(len(src)))
	stdEncoding.Encode(dst, src)
	return dst
}

func EncodeBase64String(src []byte) string {
	return stdEncoding.EncodeToString(src)
}

func EncodeBase64URLSafe(src []byte) []byte {
	dst := make([]byte, urlEncoding.EncodedLen(len(src)))
	urlEncoding.Encode(dst, src)
	return dst
}

func EncodeBase64URLSafeString(src []byte) string {
	return urlEncoding.EncodeToString(src)
}

func DecodeBase64(src []byte) ([]byte, error) {

	dst := make([]byte, stdEncoding.DecodedLen(len(src)))
	n, err := stdEncoding.Decode(dst, src)
	return dst[:n], err
}

func DecodeBase64String(s string) ([]byte, error) {
	if strings.ContainsAny(s, "+/") {
		return stdEncoding.DecodeString(s)
	}
	return urlEncoding.DecodeString(s)
}

func DecodeBase64StringWildly(s string) []byte {
	data, err := DecodeBase64String(s)
	if err != nil {
		ulogs.Panic("decode base64 string [%s] error, %v", s, err)
	}
	return data
}
