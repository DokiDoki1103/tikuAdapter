package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/pkg/errors"
)

type SearchEnncyClient struct {
}

func (in *SearchEnncyClient) getHttpClient() *resty.Client {
	return resty.New()
}

func (in *SearchEnncyClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	return nil, errors.ErrTargetNoAnswer
}
