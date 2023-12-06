package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"regexp"
)

// FillHash 填充题库的hash值
func FillHash(t *entity.Tiku) {
	reg := regexp.MustCompile("[^\\p{Han}0-9a-zA-Z]+")
	// 保留汉字、数字和字母，移除其他字符
	filteredStr := reg.ReplaceAllString(t.Question, "")

	t.QuestionHash = strutil.ShortMd5(filteredStr)
	t.Hash = strutil.Md5(t.QuestionHash + t.Options + string(t.Type))
}

// CollectAnswer 收集答案
func CollectAnswer(resp model.SearchResponse) {
	if len(resp.Answer.BestAnswer) > 0 {

		ans, _ := json.Marshal(resp.Answer.BestAnswer)
		opts, _ := json.Marshal(resp.Options)
		t := entity.Tiku{
			Question: resp.Question,
			Answer:   string(ans),
			Options:  string(opts),
			Plat:     int32(resp.Plat),
			Source:   0,
		}
		FillHash(&t)
		err := dao.Tiku.Create(&t)
		if err != nil {
			logger.SysError(fmt.Sprintf("收集答案失败 %v", err))
		}
	}
}

// CollectEmptyAnswer 收集没有搜索到的答案
func CollectEmptyAnswer(resp model.SearchRequest) {
	opts, _ := json.Marshal(resp.Options)
	t := entity.Tiku{
		Question: resp.Question,
		Answer:   "[]",
		Options:  string(opts),
		Plat:     int32(resp.Plat),
		Source:   -1,
	}
	FillHash(&t)
	err := dao.Tiku.Create(&t)
	if err != nil {
		logger.SysError(fmt.Sprintf("收集答案失败 %v", err))
	}
}
