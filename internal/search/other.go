package search

import (
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

// KV -
type KV struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// API 自定义的一些API
type API struct {
	Enable  bool   `yaml:"enable"`
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Method  string `yaml:"method"`
	Headers []KV   `yaml:"headers"`
	Body    string `yaml:"body"`
	Answer  string `yaml:"answer"`
}

func (in API) getHTTPClient() *resty.Client {
	client := resty.New().SetTimeout(5 * time.Second)
	for i := range in.Headers {
		client.SetHeader(in.Headers[i].Key, in.Headers[i].Value)
	}
	return client
}

// SearchAnswer -
func (in API) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if !in.Enable {
		return answer, nil
	}
	client := in.getHTTPClient()
	r := client.R()
	var resp *resty.Response
	if in.Method == "POST" {
		r.SetBody(replace(in.Body, req))
		resp, err = r.Post(replace(in.URL, req))
	} else {
		resp, err = r.Get(replace(in.URL, req))
	}
	if err != nil {
		return nil, err
	}

	if gjson.Valid(resp.String()) {
		ans := gjson.Get(resp.String(), in.Answer).String()
		if ans != "" {
			s := strings.Split(ans, "#")
			answer = append(answer, s)
		}
	}
	return answer, nil
}

func replace(s string, req model.SearchRequest) string {
	s = strings.Replace(s, "${question}", req.Question, -1)
	s = strings.Replace(s, "${type}", strconv.Itoa(int(req.Type)), -1)
	return s
}
