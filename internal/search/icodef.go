package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"strings"
	"time"
)

type iapiResponse struct {
	Data string `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SearchIcodefClient struct {
}

func (in *SearchIcodefClient) getHttpClient() *resty.Client {
	return resty.New().SetTimeout(5 * time.Second).SetRetryCount(2)
}

func (in *SearchIcodefClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	post, err := in.getHttpClient().R().SetFormData(map[string]string{
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
