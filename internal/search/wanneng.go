package search

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"strings"
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
		SetRetryCount(3). // 目前来看万能免费题库限流措施是4秒一次，所以做最大重试次数为3
		SetRetryWaitTime(2*time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || strings.Contains(r.String(), "已限流,正在重新请求...")
		}).
		SetRetryMaxWaitTime(10*time.Second).
		AddRetryHook(func(r *resty.Response, err error) {
			logger.SysError(fmt.Sprintf("万能免费题库触发流控限制，正在重试...%s", r.String()))
		}).
		SetHeader("Content-Type", "application/json")
}

func (in *SearchWannengClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	if in.Disable {
		return nil, errors.ErrDisable
	}
	req.Options = make([]string, 0)
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
