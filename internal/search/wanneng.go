package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/itihey/tikuAdapter/internal/model"
	"net/http"
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

func (c *searchWannengClient) SearchAnswer(req model.SearchRequest) (res model.SearchResponse, err error) {
	postData := []byte(fmt.Sprintf(`{
        "question": "%s"
    }`, req.Question))

	resp, err := http.Post("http://lyck6.cn/scriptService/api/autoFreeAnswer", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var response ApiResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return res, err
		}

		if response.Code != 0 {
			return res, errors.New(fmt.Sprintf("请求失败状态码 %d", response.Code))
		}
		// 提取answers数组的第一个元素
		if len(response.Result.Answers) > 0 {
			return model.SearchResponse{
				Question: req.Question,
				Plat:     req.Plat,
				Options:  req.Options,
				Type:     req.Type,
				Answer:   response.Result.Answers[0],
			}, nil
		}
		return res, errors.New("未找到答案")

	} else {
		return res, errors.New(fmt.Sprintf("请求失败状态码 %d", resp.StatusCode))
	}
}
