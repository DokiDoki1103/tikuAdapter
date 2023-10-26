package util

import (
	"regexp"
	"strings"
)

func formatAnswer(answers [][]string, type_ uint) {
	for i := range answers {
		answers[i] = formatSingleAnswer(answers[i], type_)
	}
}

func formatSingleAnswer(answer []string, type_ uint) []string {
	if type_ == 3 && len(answer) > 0 { // 判断题
		isTrue := regexp.MustCompile(`^(正确|是|对|√|T|ri|true|A)$`).MatchString(answer[0])
		isFalse := regexp.MustCompile(`^(错误|否|错|×|F|fa|false|B)$`).MatchString(answer[0])
		if isTrue {
			return []string{"正确"}
		} else if isFalse {
			return []string{"错误"}
		}
	} else if type_ <= 1 { // 选择题的格式化暂时还没有实现
		for i := range answer {
			answer[i] = formatString(answer[i])
		}
	}
	return answer
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

func FullWidthStrToHalfWidthStr(str string) (result string) {
	for _, charCode := range str {
		inside_code := charCode
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}

		if inside_code < 32 || inside_code > 126 {
			result += string(charCode)
		} else {
			result += string(inside_code)
		}
	}

	return result
}
