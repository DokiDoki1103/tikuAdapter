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

// GetQuestionText 保留汉字、数字和字母，移除其他字符
func GetQuestionText(q string) string {
	reg := regexp.MustCompile("[^\\p{Han}0-9a-zA-Z]+")
	return reg.ReplaceAllString(q, "")
}
