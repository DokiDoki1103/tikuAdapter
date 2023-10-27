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

type SearchIcodefClient struct {
	Disable bool
	Token   string
}

func (in *SearchIcodefClient) getHttpClient() *resty.Client {
	return resty.New().
		SetTimeout(5*time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || strings.Contains(r.String(), "触发流控限制")
		}).
		AddRetryHook(func(r *resty.Response, err error) {
			logger.SysError(fmt.Sprintf("iCodef触发流控限制，正在重试...%s", r.String()))
		}).
		SetRetryMaxWaitTime(10*time.Second).
		SetHeader("Authorization", in.Token)
}

func (in *SearchIcodefClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	post, err := in.getHttpClient().R().
		SetFormData(map[string]string{
			"question": req.Question,
		}).Post("https://cx.icodef.com/wyn-nb?v=4")
	if err != nil {
		return nil, errors.ErrTargetServerError
	}

	var response iapiResponse
	err = json.Unmarshal(post.Body(), &response)
	if err != nil {
		return nil, errors.ErrTargetServerError
	}

	if response.Code != 1 {
		return nil, errors.ErrTargetNoAnswer
	}

	ans := strings.Split(response.Data, "#")

	return [][]string{ans}, nil

}
