package search

import (
	"encoding/base64"
	"github.com/antlabs/strsim"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

// BuguakeClient 不挂科网页版题库
type BuguakeClient struct {
	Disable bool
}

func (in *BuguakeClient) getHTTPClient() *resty.Client {
	return resty.New()
}

// SearchAnswer 搜索答案
func (in *BuguakeClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if in.Disable {
		return answer, nil
	}

	client := in.getHTTPClient()
	resp, err := client.R().
		SetQueryParam("query", req.Question).
		SetQueryParam("rn", "10").
		SetQueryParam("pn", "0").
		Get("https://easylearn.baidu.com/edu-web-go/bgk/searchlist")
	if err != nil {
		return nil, errors.ErrTargetServerError
	}

	list := gjson.Get(resp.String(), "data.list").Array()

	for _, value := range list {
		dec := decrypt(value.Get("bdjson").Str, value.Get("actk").Str)

		s := strsim.Compare(req.Question, gjson.Get(dec, "que_stem.0.c.0.c").String())
		if s < 0.8 { // 百度题库太多不相似然后随机返回的，低于0.8的不要
			continue
		}
		as := make([]string, 0)
		// 优先读取选项
		for _, val := range gjson.Get(dec, "que_options.0.#(yes=true).ret.0.c.0.c").Array() {
			as = append(as, val.Str)
		}
		if len(as) > 0 {
			answer = append(answer, as)
			continue
		}

		ans := gjson.Get(dec, "que_answer.0")
		if ans.IsObject() {
			for _, v := range ans.Get("c.#.c").Array() {
				as = append(as, v.Str)
			}
			if len(as) > 0 {
				answer = append(answer, as)
				continue
			}
		} else if ans.IsArray() {
			// options := gjson.Get(dec, "que_stem.1.c.0.c").String()
			// 需要从选项中提取 比较麻烦，暂时没做
		}

	}
	return answer, nil

}

func decrypt(t, r string) string {
	n := "34cab29ef956d78afd"
	e := n + r

	// 对输入字符串t进行Base64解码
	tDecoded, _ := base64.StdEncoding.DecodeString(t)

	var o int
	i := make([]byte, len(tDecoded))

	for c := 0; c < len(tDecoded); c++ {
		o = o % len(e)
		a := int(tDecoded[c])
		f := int(e[o])
		s := byte(a ^ f)
		i[c] = s
		o++
	}

	var l string

	for p := 0; p < len(i)-1; p++ {
		d := int(i[p]) ^ int(i[p+1])
		l += string(rune(d))
		p++
	}

	// 对字符串l进行URI解码并替换+
	h, _ := url.QueryUnescape(l)
	h = strings.ReplaceAll(h, "+", " ")
	return h
}
