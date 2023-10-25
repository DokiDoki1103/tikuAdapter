package util

import (
	"sort"
	"strings"
)

// SearchRightAnswer 此方法还有巨大的优化空间
// 1: 统一格式化每个答案的内容后再比对
// 2: 字符串相似度匹配-模糊搜索
// 3: 如果类型是判断题那么 对|正确|True 应该统一格式化为正确
// 4: 如果用户给了 options 直接计算出答案的角标进行返回 [0,1,2]
func SearchRightAnswer(answers [][]string) []string {
	answerCount := make(map[string]int)
	sep := "**=====^_^======^_^=====**"
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
