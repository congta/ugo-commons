package unumbers

import (
	"strconv"

	"github.com/congta/ugo-commons/commons-logging/ulogs"
)

// ParseInt get int, return error if input is not a number
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseIntWildly get int, panic if input is not a number
func ParseIntWildly(s string) int {
	i, err := ParseInt(s)
	if err != nil {
		ulogs.Panic("convert non-number string %s to int error. %v", s, err)
	}
	return i
}

// ToInt get int, return default value if input is not a number
func ToInt(s string, defaultValue int) int {
	i, err := ParseInt(s)
	if err != nil {
		return defaultValue
	}
	return i
}
