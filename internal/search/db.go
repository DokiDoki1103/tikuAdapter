package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/registry/m"
	"github.com/itihey/tikuAdapter/pkg/model"
)

// DB mysql 或者sqlite3
type DB struct{}

func (in *DB) getHTTPClient() *resty.Client {
	panic("implement me")
}

// SearchAnswer 搜索答案
func (in *DB) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	questionHash := strutil.ShortMd5(req.Question)
	tiku := m.GetManager().Query.Tiku
	find, err := tiku.Where(tiku.QuestionHash.Eq(questionHash)).Find()
	if err != nil {
		return nil, err
	}
	for i := range find {
		var answers []string // 最后所有的答案的二维数组
		err := json.Unmarshal([]byte(find[i].Answer), &answers)
		if err != nil {
			continue
		}
		answer = append(answer, answers)
	}
	return answer, nil
}
