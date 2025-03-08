package search

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
)

type TikuhaiResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type TikuhaiResult struct {
	Answer []string `json:"answer"`
	Num    int      `json:"num"`
	Usenum int      `json:"usenum"`
}

// TikuhaiClient 题海题库客户端
type TikuhaiClient struct {
	Enable bool
	Token  string
}

func (in *TikuhaiClient) getHTTPClient() *resty.Client {
	return resty.New().
		SetTimeout(5 * time.Second).
		SetRetryCount(1).
		SetRetryWaitTime(2 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() >= 500
		}).
		AddRetryHook(func(r *resty.Response, err error) {
			logger.SysError(fmt.Sprintf("题库海题库触发重试，状态码：%d，响应：%s", r.StatusCode(), r.String()))
		}).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   "tikuhaiAdapter/0.1.0",
			"v":            "0.1.0",
		})
}

func (in *TikuhaiClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if !in.Enable {
		return answer, nil
	}

	// 构造请求参数
	body := map[string]interface{}{
		"question":     req.Question,
		"options":      req.Options,
		"type":         req.Type,
		"key":          in.Token,
		"questionData": "",
	}

	resp, err := in.getHTTPClient().R().
		SetBody(body).
		Post("https://api.tikuhai.com/search")

	if err != nil || resp.StatusCode() != 200 {
		return nil, errors.ErrTargetServerError
	}

	var res TikuhaiResponse

	if err := json.Unmarshal(resp.Body(), &res); err != nil {

		return nil, errors.ErrParserJSON
	}

	switch {
	case res.Code == 200:
		var result TikuhaiResult
		if err := json.Unmarshal(res.Data, &result); err != nil {
			return nil, errors.ErrParserJSON
		}
		if len(result.Answer) > 0 {
			return [][]string{result.Answer}, nil
		}
		//200 有答案应该不会出现这种情况
		return nil, errors.ErrTargetNoAnswer

	case res.Code == -1 && strings.Contains(res.Msg, "有答案"):
		return nil, errors.ErrTokenRequired

	case res.Code == -1:
		return nil, errors.ErrTargetNoAnswer

	default:
		return nil, errors.ErrTargetServerError
	}
}
