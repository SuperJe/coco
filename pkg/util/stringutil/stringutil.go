package stringutil

import "strings"

// CapFirstLetter 将第一个字母转成大写
func CapFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
