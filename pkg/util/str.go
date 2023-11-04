package util

import (
	"unicode"
)

// IsAlpha 判断字符串是否全是字母
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
