package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/model"
	"strings"
	"time"
)

// EnncyClient enncy题库
type EnncyClient struct {
	Enable bool   // 是否禁用
	Token  string // token
}

// getHTTPClient 获取HTTP客户端
func (in *EnncyClient) getHTTPClient() *resty.Client {
	return resty.New().SetTimeout(5 * time.Second)
}

// SearchAnswer 搜索答案
func (in *EnncyClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if !in.Enable || in.Token == "" {
		return answer, nil
	}
	post, err := in.getHTTPClient().R().
		SetQueryParam("token", in.Token).
		SetQueryParam("title", req.Question).
		Get("https://tk.enncy.cn/query")
	if err != nil || post.StatusCode() != 200 {
		return nil, errors.ErrTargetServerError
	}
	var response enncyResponse
	err = json.Unmarshal(post.Body(), &response)
	if err != nil {
		return nil, errors.ErrTargetServerError
	}
	if response.Code == 1 {
		ans := strings.Split(response.Data.Answer, "#")
		return [][]string{ans}, nil
	}
	return nil, errors.ErrTargetNoAnswer
}

type enncyData struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Times    int    `json:"times"`
}

type enncyResponse struct {
	Code    int       `json:"code"`
	Data    enncyData `json:"data"`
	Message string    `json:"message"`
}
