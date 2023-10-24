package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
	"time"
)

type searchEnncyClient struct {
	s search
}

func (in *searchEnncyClient) getHttpClient() *resty.Client {
	return resty.New().SetTimeout(5 * time.Second)
}

func (in *searchEnncyClient) SearchAnswer(req model.SearchRequest) (res model.SearchResponse, err error) {
	panic("impl me")
}
