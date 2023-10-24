package model

// SearchRequest 搜题的请求题
type SearchRequest struct {
	Plat     uint     `json:"plat"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Type     uint     `json:"type"`
}

// SearchResponse 搜题响应体
type SearchResponse struct {
	Plat       uint       `json:"plat"`
	Question   string     `json:"question"`
	Options    []string   `json:"options"`
	Type       uint       `json:"type"`
	Answer     []string   `json:"answer"`     //最相似，最可能的答案->即moreAnswer中出现最多的答案
	MoreAnswer [][]string `json:"moreAnswer"` // 所有接口聚合的答案
}
