package uconv

import (
	"github.com/duke-git/lancet/v2/convertor"
	"math"
)

func IntPtr(i int) *int {
	return &i
}

func Uint64ToInt64(n uint64) int64 {
	return int64(n & math.MaxInt64)
}

var (
	ToInt = convertor.ToInt
)

func ToInt32(n any) (int32, error) {
	i, err := ToInt(n)
	return int32(i & 0xffff), err
}
