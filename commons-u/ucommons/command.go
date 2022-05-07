package ucommons

import "strings"

func GetArgsMap(args []string) map[string]string {
	result := make(map[string]string)
	num := len(args)
	for i := 0; i < num; i++ {
		if !strings.HasPrefix(args[i], "-") {
			continue
		}
		key := args[i]
		value := ""
		eqIdx := strings.Index(key, "=")
		if eqIdx > 1 {
			// if eqIdx == 1, the input should looks like "-= 0", regard "-=" as key
			value = key[eqIdx+1:]
			key = key[:eqIdx] // change key at last
		}
		for i+1 < num && !strings.HasPrefix(args[i+1], "-") {
			if value == "" {
				value += args[i+1]
			} else {
				value += " " + args[i+1]
			}
			i++
		}
		result[key] = value
	}
	return result
}
