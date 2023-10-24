package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"time"
)

type searchWannengClient struct {
	s search
}

type ApiResponse struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Result    Result `json:"result"`
}

type Result struct {
	Group   string     `json:"group"`
	Num     int        `json:"num"`
	Answers [][]string `json:"answers"`
	Success bool       `json:"success"`
}

func (in *searchWannengClient) getHttpClient() *resty.Client {
	return resty.New().SetTimeout(5 * time.Second)
}

func (in *searchWannengClient) SearchAnswer(req model.SearchRequest) (res model.SearchResponse, err error) {
	client := in.getHttpClient()

	data, _ := json.Marshal(req)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(data)).
		Post("http://lyck6.cn/scriptService/api/autoFreeAnswer")

	if err != nil || resp.StatusCode() != 200 {
		return res, errors.ErrTargetServerError
	}

	var response ApiResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return model.SearchResponse{}, errors.ErrTargetServerError
	}
	// 不等于0就是请求失败
	if response.Code != 0 {
		if response.Code == 429 {
			return res, errors.ErrTargetApiFlow
		}
		return res, errors.ErrTargetServerError
	}
	if len(response.Result.Answers) == 0 {
		return res, errors.ErrTargetNoAnswer
	}
	return model.SearchResponse{
		Question:   req.Question,
		Plat:       req.Plat,
		Options:    req.Options,
		Type:       req.Type,
		Answer:     response.Result.Answers[0],
		MoreAnswer: response.Result.Answers,
	}, nil
}
