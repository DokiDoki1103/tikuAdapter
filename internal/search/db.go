package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/pkg/model"
	"log"
	"sort"
	"strconv"
)

// DB mysql 或者sqlite3
type dBSearch struct{}

var defaultDBSearch = &dBSearch{}

// GetDBSearch 获取DB搜索实例
func GetDBSearch() Search {
	return defaultDBSearch
}
func (in *dBSearch) getHTTPClient() *resty.Client {
	panic("implement me")
}

// SearchAnswer 搜索答案
func (in *dBSearch) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	// 将用户传递来的选项进行排序
	sortOptions := make([]string, len(req.Options))
	copy(sortOptions, req.Options)
	sort.Strings(sortOptions)
	sortOptionsStr, err1 := json.Marshal(sortOptions)
	if err1 != nil {
		sortOptionsStr = []byte("[]")
	}
	// 生成hash值
	Hash := strutil.Md5(req.Question + string(sortOptionsStr) + strconv.Itoa(req.Type) + strconv.Itoa(req.Plat))
	tiku := dao.Tiku
	find, err := tiku.Where(tiku.Hash.Eq(Hash)).Find()
	if err != nil {
		return nil, err
	}
	// 如果数据库中没有extra那么自动补全他
	if len(find) == 1 && find[0].Extra == "" {
		_, err1 := tiku.Where(tiku.ID.Eq(find[0].ID)).Update(tiku.Extra, req.Extra)
		if err1 != nil {
			log.Println("更新extra失败", err1)
		}
	}
	for i := range find {
		var answers []string // 最后所有的答案的二维数组
		err := json.Unmarshal([]byte(find[i].Answer), &answers)
		if err != nil {
			continue
		}
		if len(answers) > 0 {
			answer = append(answer, answers)
		}
	}
	return answer, nil
}
