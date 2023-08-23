package ustrings

import "strings"

// SubstringAfterLast substring after last separator, return raw str if separator is not exist.
// (the action is different from strutil.AfterLast in lancet when separator is not exist)
func SubstringAfterLast(str, separator string) string {
	if str == "" || separator == "" {
		return str
	}
	idx := strings.LastIndex(str, separator)
	if idx < 0 {
		return str
	} else {
		return str[idx+len(separator):]
	}
}
