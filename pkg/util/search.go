package util

import (
	"github.com/gookit/goutil/arrutil"
	"github.com/itihey/tikuAdapter/internal/model"
	"sort"
	"strings"
)

var sep = "**=====^_^======^_^=====**" // 用于分割答案的分隔符

func FillAnswerResponse(answers [][]string, req *model.SearchRequest) model.SearchResponse {
	resp := model.SearchResponse{
		MoreAnswer: answers,
		Question:   req.Question,
		Options:    req.Options,
		Type:       req.Type,
		Plat:       req.Plat,
	}
	formatAnswer(answers, req.Type) // 先把答案统一格式化
	for i := range resp.Options {
		resp.Options[i] = formatString(resp.Options[i])
	}

	if resp.AnswerIndex == nil {
		resp.AnswerIndex = make([]int, 0)
	}

	if req.Options == nil || len(req.Options) == 0 { // 用户没有传选项，那么只能返回出现次数最多的答案。
		resp.Answer = SearchRightAnswer(answers, req)
	} else {
		answerCount := make(map[string]int)
		for i := range answers {
			ans := arrutil.Intersects(req.Options, answers[i], arrutil.StringEqualsComparer)
			if uint(len(ans)) > resp.Type {
				answerCount[strings.Join(ans, sep)]++
			}
		}
		resp.Answer = getMaxCountAnswer(answerCount)

		// if len(resp.Answer) == 0 {
		// 	for i := range answers {
		// 		match := strsim.FindBestMatch(answers[i][0], req.Options)
		// 		resp.AnswerIndex = append(resp.AnswerIndex, match.BestIndex)
		// 	}
		// 	resp.Answer = []string{req.Options[resp.AnswerIndex[0]]}
		// }
		resp.AnswerIndex = findIndices(resp.Answer, req.Options)
	}

	return resp
}

// SearchRightAnswer 此方法还有巨大的优化空间
func SearchRightAnswer(answers [][]string, s *model.SearchRequest) []string {
	answerCount := make(map[string]int)
	for _, answer := range answers {
		sort.Strings(answer)
		sortedAnswer := strings.Join(answer, sep)
		answerCount[sortedAnswer]++
	}
	return getMaxCountAnswer(answerCount)
}

func getMaxCountAnswer(answerCount map[string]int) []string {
	maxCount := 0
	var correctAnswers []string

	for answer, count := range answerCount {
		newAnswers := strings.Split(answer, sep)
		if count > maxCount || len(newAnswers) > len(correctAnswers) {
			maxCount = count
			correctAnswers = newAnswers
		}
	}
	return correctAnswers
}

func findIndices(answers []string, options []string) []int {
	indices := make([]int, 0)
	for _, answer := range answers {
		for i, option := range options {
			if option == answer {
				indices = append(indices, i)
				break
			}
		}
	}
	return indices
}
