package model

// SearchRequest 搜题的请求题
type SearchRequest struct {
	Plat     string   `json:"plat"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Type     uint     `json:"type"`
}

// SearchResponse 搜题响应体
type SearchResponse struct {
	Plat     string   `json:"plat"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Type     uint     `json:"type"`
	Answer   []string `json:"answer"`
}
