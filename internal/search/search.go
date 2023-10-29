package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/model"
)

type SearchClient struct {
	Wanneng *SearchWannengClient
	Enncy   *SearchEnncyClient
	Icodef  *SearchIcodefClient
}

type Search interface {
	getHttpClient() *resty.Client
	SearchAnswer(req model.SearchRequest) (answer [][]string, err error)
}
