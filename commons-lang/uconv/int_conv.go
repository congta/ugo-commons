package uconv

import "math"

func IntPtr(i int) *int {
	return &i
}

func Uint64ToInt64(n uint64) int64 {
	return int64(n & math.MaxInt64)
}
