package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
)

type SearchClient struct {
	HttpClient *resty.Client
	Wanneng    searchWannengClient
	Enncy      searchEnncyClient
}

// Search 所有的请求外部的题库接口都需要自行实现此接口
type search interface {
	getHttpClient() *resty.Client
	SearchAnswer(req model.SearchRequest) (res model.SearchResponse, err error)
}
