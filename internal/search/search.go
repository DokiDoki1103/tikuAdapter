package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
)

type SearchClient struct {
	Wanneng *SearchWannengClient
	Enncy   *SearchEnncyClient
	Icodef  *SearchIcodefClient
}

// Search 所有的请求外部的题库接口都需要自行实现此接口
type search interface {
	getHttpClient() *resty.Client
	SearchAnswer(req model.SearchRequest) (answer [][]string, err error)
}
