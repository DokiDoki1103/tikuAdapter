package model

// SearchRequest 搜题的请求体
type SearchRequest struct {
	QID      string   `json:"qid"`
	Plat     uint     `json:"plat"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Type     uint     `json:"type"`

	Extra string `json:"extra"`
}

// SearchResponse 搜题响应体
type SearchResponse struct {
	Plat     uint     `json:"plat"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Type     uint     `json:"type"`
	Answer   Answer   `json:"answer"` // 最相似，最可能的答案->即moreAnswer中出现最多的答案
}

// Answer 答案 该实体只会增加而不会减少
type Answer struct {
	AnswerKey     []string `json:"answerKey"`     // 如果传了选项可以格式化为ABCD
	AnswerKeyText string   `json:"answerKeyText"` // 答案ABCD文本 分隔符默认是 #

	AnswerIndex []int  `json:"answerIndex"` // 如果用户传了options 将会自动帮用户计算出答案的角标，依据Answer来计算的
	AnswerText  string `json:"answerText"`  // 答案文本 分隔符默认是 #

	BestAnswer []string   `json:"bestAnswer"` // 最有可能的答案
	AllAnswer  [][]string `json:"allAnswer"`  // 所有接口聚合的答案
}

// Question 最简单的通用问题答案模型
type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Type     uint     `json:"type"`
	Answer   []string `json:"answer"`
}
