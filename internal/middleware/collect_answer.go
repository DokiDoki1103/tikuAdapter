package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/itihey/tikuAdapter/pkg/util"
)

// FillHash 填充题库的hash值
func FillHash(t *entity.Tiku) {
	t.Question = util.FormatString(t.Question)

	questionText := util.GetQuestionText(t.Question)
	t.QuestionText = questionText
	t.QuestionHash = strutil.ShortMd5(questionText)
	t.Hash = strutil.Md5(t.QuestionHash + t.Options + string(t.Type) + t.Extra)

	if t.Answer == "" {
		t.Answer = "[]"
	} else if t.Options == "" {
		t.Options = "[]"
	}
}

// CollectEmptyAnswer 收集没有搜索到的答案
func CollectEmptyAnswer(resp model.SearchRequest) {
	opts, _ := json.Marshal(resp.Options)
	t := entity.Tiku{
		Type:     int32(resp.Type),
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
