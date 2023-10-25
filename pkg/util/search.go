package util

import (
	"regexp"
	"sort"
	"strings"
)

var sep = "**=====^_^======^_^=====**" // 用于分割答案的分隔符

// SearchRightAnswer 此方法还有巨大的优化空间
// 1: 统一格式化每个答案的内容后再比对  ✅
// 2: 字符串相似度匹配-模糊搜索
// 3: 如果类型是判断题那么 对|正确|True 应该统一格式化为正确 ✅
// 4: 如果用户给了 options 直接计算出答案的角标进行返回 [0,1,2]
func SearchRightAnswer(answers [][]string, type_ uint) []string {
	formatAnswer(answers, type_)
	answerCount := make(map[string]int)
	for _, answer := range answers {
		sort.Strings(answer)
		sortedAnswer := strings.Join(answer, sep)
		answerCount[sortedAnswer]++
	}

	maxCount := 0
	var correctAnswers []string

	for answer, count := range answerCount {
		if count > maxCount {
			maxCount = count
			correctAnswers = strings.Split(answer, sep)
		}
	}
	return correctAnswers
}

func formatString(src string) string {
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

func formatAnswer(answers [][]string, type_ uint) {
	for i := range answers {
		answers[i] = formatSingleAnswer(answers[i], type_)
	}
}
