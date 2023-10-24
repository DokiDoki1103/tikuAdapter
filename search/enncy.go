package search

import "github.com/itihey/tikuAdapter/model"

type searchEnncyClient struct {
	s search
}

func (c *searchEnncyClient) SearchAnswer(req model.SearchRequest) (res model.SearchResponse, err error) {
	panic("impl me")
}
