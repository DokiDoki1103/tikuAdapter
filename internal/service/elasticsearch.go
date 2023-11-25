package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/itihey/tikuAdapter/internal/dao"
	"strconv"
)

// Elasticsearch Elasticsearch接口
type Elasticsearch interface {
	// Search 搜索
	Search(index string, query string, from int, size int) (interface{}, error)
	// Create 创建
	Create(index string, id string, data interface{}) error
	// Update 更新
	Update(index string, id string, data interface{}) error
	// Delete 删除
	Delete(index string, id string) error
}

type elasticSearch struct {
	es *elasticsearch.Client
}

var defaultElasticSearch Elasticsearch

// NewElasticsearchClient 创建Elasticsearch客户端
func NewElasticsearchClient(add []string) (Elasticsearch, error) {
	cfg := elasticsearch.Config{
		Addresses: add,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	defaultElasticSearch = &elasticSearch{es: es}
	return defaultElasticSearch, nil
}

// Search 搜索
func (e *elasticSearch) Search(index string, query string, from int, size int) (interface{}, error) {
	var buf bytes.Buffer
	queryJSON := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"word": query,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(queryJSON); err != nil {
		return nil, err
	}

	res, err := e.es.Search(
		e.es.Search.WithContext(context.Background()),
		e.es.Search.WithIndex(index),
		e.es.Search.WithBody(&buf),
		e.es.Search.WithTrackTotalHits(true),
		e.es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var searchResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		return nil, err
	}

	return searchResult, nil
}

// Create 创建
func (e *elasticSearch) Create(index string, id string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res, err := e.es.Index(index, bytes.NewReader(body), e.es.Index.WithDocumentID(id))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil // 这里需要根据实际情况处理错误
	}

	return nil
}

// Update 更新
func (e *elasticSearch) Update(index string, id string, data interface{}) error {
	body, err := json.Marshal(map[string]interface{}{
		"word": data,
	})
	if err != nil {
		return err
	}

	res, err := e.es.Update(index, id, bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil // 这里需要根据实际情况处理错误
	}

	return nil
}

// Delete 删除
func (e *elasticSearch) Delete(index string, id string) error {
	res, err := e.es.Delete(index, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil // 这里需要根据实际情况处理错误
	}

	return nil
}

// SyncElasticsearch 同步Elasticsearch
func SyncElasticsearch() {
	find, err := dao.Tiku.Find()
	if err != nil {
		return
	}
	for _, v := range find {
		err := defaultElasticSearch.Create("tiku", strconv.Itoa(int(v.ID)), v)
		if err != nil {
			fmt.Println("es create err: ", err)
			return
		}
	}
}
