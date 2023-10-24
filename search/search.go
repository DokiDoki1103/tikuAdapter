package search

import "github.com/itihey/tikuAdapter/model"

type SearchClient struct {
	Wanneng searchWannengClient
	Enncy   searchEnncyClient
}

// Search 所有的请求外部的题库接口都需要自行实现此接口
type search interface {
	SearchAnswer(req model.SearchRequest) (res model.SearchResponse, err error)
}
