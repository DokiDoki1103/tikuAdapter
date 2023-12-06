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

type iapiResponse struct {
	Data string `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// IcodefClient icodef题库
type IcodefClient struct {
	Enable bool
	Token  string
}

// getHTTPClient 获取HTTP客户端
func (in *IcodefClient) getHTTPClient() *resty.Client {
	return resty.New().
		SetTimeout(5*time.Second).
		SetRetryCount(1).
		SetRetryWaitTime(time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || (strings.Contains(r.String(), "触发流控限制") && !strings.Contains(r.String(), "IP超出每日限额"))
		}).
		AddRetryHook(func(r *resty.Response, err error) {
			logger.SysError(fmt.Sprintf("iCodef触发流控限制，正在重试...%s", r.String()))
		}).
		SetRetryMaxWaitTime(10*time.Second).
		SetHeader("Authorization", in.Token)
}

// SearchAnswer 搜索答案
func (in *IcodefClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if !in.Enable {
		return answer, nil
	}

	post, err := in.getHTTPClient().R().
		SetFormData(map[string]string{
			"question": req.Question,
		}).Post("https://cx.icodef.com/wyn-nb?v=4")
	if err != nil {
		return nil, errors.ErrTargetServerError
	}

	var r iapiResponse
	err = json.Unmarshal(post.Body(), &r)
	if err != nil {
		return nil, errors.ErrTargetServerError
	}
	if r.Code != 1 {
		return nil, errors.ErrTargetNoAnswer
	}
	ans := strings.Split(r.Data, "#")
	return [][]string{ans}, nil
}
