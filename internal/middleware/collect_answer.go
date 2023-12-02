package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
)

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
		t.QuestionHash = strutil.ShortMd5(t.Question)
		t.Hash = strutil.Md5(t.QuestionHash + t.Options + string(t.Type))
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
	t.QuestionHash = strutil.ShortMd5(t.Question)
	t.Hash = strutil.Md5(t.QuestionHash + t.Options + string(t.Type))
	err := dao.Tiku.Create(&t)
	if err != nil {
		logger.SysError(fmt.Sprintf("收集答案失败 %v", err))
	}
}
