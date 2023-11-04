package util

import (
	"github.com/gookit/goutil/strutil"
	"regexp"
	"strings"
)

func formatAnswer(answers [][]string, questionType uint) {
	for i := range answers {
		answers[i] = formatSingleAnswer(answers[i], questionType)
	}
}

func formatSingleAnswer(answer []string, questionType uint) []string {
	if questionType == 3 && len(answer) > 0 { // 判断题
		isTrue := regexp.MustCompile(`^(正确|是|对|√|T|ri|true|A)$`).MatchString(answer[0])
		isFalse := regexp.MustCompile(`^(错误|否|错|×|F|fa|false|B)$`).MatchString(answer[0])
		if isTrue {
			return []string{"正确"}
		} else if isFalse {
			return []string{"错误"}
		}
	} else if questionType <= 1 { // 选择题的格式化暂时还没有实现
		for i := range answer {
			answer[i] = formatString(answer[i])
		}
	}
	return answer
}

// FormatOptions 格式化选项
func FormatOptions(options []string) []string {
	var formattedOptions []string
	re := regexp.MustCompile(`^[A-Z][.．:：、]\s?`)
	for _, option := range options {
		formattedOption := strutil.Trim(re.ReplaceAllString(option, ""), " \t\n\r")
		formattedOptions = append(formattedOptions, formattedOption)
	}
	return formattedOptions
}

func formatString(src string) string {
	// 全角转半角
	src = FullWidthStrToHalfWidthStr(src)
	// 中文常见符号转英文
	src = strings.ReplaceAll(src, "“", `"`)
	src = strings.ReplaceAll(src, "”", `"`)
	src = strings.ReplaceAll(src, "‘", "'")
	src = strings.ReplaceAll(src, "’", "'")
	src = strings.ReplaceAll(src, "。", ".")

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
