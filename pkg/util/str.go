package util

import (
	"regexp"
)

// IsAlpha 判断字符串是否全是大写字母
func IsAlpha(s string) bool {
	m, err := regexp.MatchString(`^[A-Z]+$`, s)
	if err != nil {
		return false
	}
	return m
}
