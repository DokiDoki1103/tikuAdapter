package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/model"
	"time"
)

// ChatGptClient chatgpt搜题客户
type ChatGptClient struct {
	Disable bool   // 是否禁用
	Token   string // token
	Host    string // openAi host
}

// getHTTPClient 获取HTTP客户端
func (in *ChatGptClient) getHTTPClient() *resty.Client {
	return resty.New().SetTimeout(5 * time.Second)
}

// SearchAnswer 暂时未实现
func (in *ChatGptClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	return nil, errors.ErrTargetNoAnswer
}
