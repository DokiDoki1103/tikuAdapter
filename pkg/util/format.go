package util

import (
	"github.com/gookit/goutil/strutil"
	"regexp"
	"strings"
)

// IsTrue 判断字符串是否是正确
func IsTrue(s string) bool {
	return regexp.MustCompile(`^(正确|是|对|√|T|ri|true|A)$`).MatchString(s)
}

// IsFalse 判断字符串是否是错误
func IsFalse(s string) bool {
	return regexp.MustCompile(`^(错误|否|错|×|F|fa|false|B)$`).MatchString(s)
}

// FormatOptions 格式化选项
func FormatOptions(options []string, questionType int) []string {
	if options == nil || len(options) == 0 {
		return []string{}
	}

	options = formatOptions(options)
	if questionType == 3 && len(options) > 0 { // 判断题
		for i := range options {
			if IsTrue(options[i]) {
				options[i] = "正确"
			} else if IsFalse(options[i]) {
				options[i] = "错误"
			}
		}
	} else if questionType <= 1 { // 选择题的格式化暂时还没有实现
		for i := range options {
			options[i] = FormatString(options[i])
		}
	}
	return options
}

func formatOptions(options []string) []string {
	var formattedOptions []string
	re := regexp.MustCompile(`^[A-Z][.．:：、]\s?`)
	for _, option := range options {
		formattedOption := strutil.Trim(re.ReplaceAllString(option, ""), " \t\n\r")
		formattedOptions = append(formattedOptions, formattedOption)
	}
	return formattedOptions
}

// FormatString 格式化字符串
func FormatString(src string) string {
	// 全角转半角
	src = FullWidthStrToHalfWidthStr(src)
	// 中文常见符号转英文
	src = strings.ReplaceAll(src, "“", `"`)
	src = strings.ReplaceAll(src, "”", `"`)
	src = strings.ReplaceAll(src, "‘", "'")
	src = strings.ReplaceAll(src, "’", "'")
	src = strings.ReplaceAll(src, "。", ".")
	src = strings.ReplaceAll(src, "&nbsp;", " ")

	// 去除末尾的常见字符
	src = strings.TrimRightFunc(src, func(r rune) bool {
		return strings.ContainsRune(",.?:!;", r)
	})

	return strings.TrimSpace(src)
}

// FullWidthStrToHalfWidthStr 全角转半角
func FullWidthStrToHalfWidthStr(str string) (result string) {
	for _, charCode := range str {
		insideCcode := charCode
		if insideCcode == 12288 {
			insideCcode = 32
		} else {
			insideCcode -= 65248
		}

		if insideCcode < 32 || insideCcode > 126 {
			result += string(charCode)
		} else {
			result += string(insideCcode)
		}
	}

	return result
}
