package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"log"
	"time"
)

type result struct {
	Group   string     `json:"group"`
	Num     int        `json:"num"`
	Answers [][]string `json:"answers"`
	Success bool       `json:"success"`
}

type wapiResponse struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Result    result `json:"result"`
}

type SearchWannengClient struct {
	Disable bool
	Token   string
}

func (in *SearchWannengClient) getHttpClient() *resty.Client {
	return resty.New().
		SetTimeout(5*time.Second).
		SetRetryCount(3).
		SetHeader("Content-Type", "application/json")
}

func (in *SearchWannengClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	log.Println(in.Disable)
	if in.Disable {
		return nil, errors.ErrDisable
	}
	data, _ := json.Marshal(req)

	url := "http://lyck6.cn/scriptService/api/autoFreeAnswer"
	if in.Token != "" && len(in.Token) == 10 {
		url = "http://lyck6.cn/scriptService/api/autoAnswer/" + in.Token
	}
	resp, err := in.getHttpClient().R().
		SetBody(string(data)).
		Post(url)

	if err != nil || resp.StatusCode() != 200 {
		return nil, errors.ErrTargetServerError
	}

	var response wapiResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, errors.ErrTargetServerError
	}
	// 不等于0就是请求失败
	if response.Code != 0 {
		if response.Code == 429 {
			return nil, errors.ErrTargetApiFlow
		}
		return nil, errors.ErrTargetServerError
	}
	if len(response.Result.Answers) == 0 {
		return nil, errors.ErrTargetNoAnswer
	}
	return response.Result.Answers, nil
}
