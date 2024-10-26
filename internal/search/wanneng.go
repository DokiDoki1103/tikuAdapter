package search

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"strconv"
	"strings"
	"time"
)

type result struct {
	Group   string        `json:"group"`
	Num     int           `json:"num"`
	Answers []interface{} `json:"answers"`
	Success bool          `json:"success"`
}

type wapiResponse struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Result    result `json:"result"`
}

// WannengClient icodef题库
type WannengClient struct {
	Enable bool
	Token  string
}

// getHTTPClient 获取HTTP客户端
func (in *WannengClient) getHTTPClient() *resty.Client {
	return resty.New().
		SetTimeout(5*time.Second).
		SetRetryCount(1).
		SetRetryWaitTime(1*time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || strings.Contains(r.String(), "已限流,正在重新请求...")
		}).
		SetRetryMaxWaitTime(10*time.Second).
		AddRetryHook(func(r *resty.Response, err error) {
			logger.SysError(fmt.Sprintf("万能题库触发流控限制:%d,正在重试...%s", r.StatusCode(), r.String()))
		}).
		SetHeader("Content-Type", "application/json")
}

// SearchAnswer 搜索答案
func (in *WannengClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if !in.Enable {
		return answer, nil
	}

	url := "http://lyck6.cn/scriptService/api/autoFreeAnswer"
	if in.Token != "" && len(in.Token) == 10 {
		url = "http://lyck6.cn/scriptService/api/autoAnswer/" + in.Token
	}
	resp, err := in.getHTTPClient().R().
		SetBody(req).
		SetHeader("plat", strconv.Itoa(req.Plat)).
		Post(url)
	if err != nil || resp.StatusCode() != 200 {
		return nil, errors.ErrTargetServerError
	}

	var res wapiResponse
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, errors.ErrTargetServerError
	}
	// 不等于0就是请求失败
	if res.Code != 0 {
		if res.Code == 429 {
			return nil, errors.ErrTargetAPIFlow
		}
		return nil, errors.ErrTargetServerError
	}
	if res.Result.Success {
		var as []string
		for _, v := range res.Result.Answers {
			as = append(as, req.Options[(int)(v.(float64))])
		}
		return [][]string{as, as, as, as, as, as, as, as, as, as}, nil
	}

	// 万能题库返回的是一个二维数组
	for _, ans := range res.Result.Answers {
		var innerArray []string
		for _, val := range ans.([]interface{}) {
			s := val.(string)
			if len(s) > 0 {
				innerArray = append(innerArray, s)
			}
		}

		if len(innerArray) > 0 {
			answer = append(answer, innerArray)
		}
	}
	fmt.Println(res.Result.Answers)
	return answer, nil
}
