package uints

import (
	"congta.com/ugo-commons/commons-logging/ulogs"
	"strconv"
)

func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseIntWildly(s string) int {
	i, err := ParseInt(s)
	if err != nil {
		ulogs.Panic("convert non-number string %s to int error. %v", s, err)
	}
	return i
}
