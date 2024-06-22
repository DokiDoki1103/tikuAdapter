package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/internal/registry/manager"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/itihey/tikuAdapter/pkg/util"
	"os"
)

// FillHash 填充题库的hash值
func FillHash(t *entity.Tiku) {
	t.Question = util.FormatString(t.Question)

	questionText := util.GetQuestionText(t.Question)
	t.QuestionText = questionText
	t.QuestionHash = strutil.ShortMd5(questionText)
	t.Hash = strutil.Md5(t.QuestionHash + t.Options + string(t.Type))

	if t.Answer == "" {
		t.Answer = "[]"
	} else if t.Options == "" {
		t.Options = "[]"
	}
}

// CollectAnswer 收集没有搜索到的答案
func CollectAnswer(resp model.SearchResponse, extra string) {
	opts, _ := json.Marshal(resp.Options)
	ans := "[]"
	if len(resp.Answer.AnswerKey) > 0 {
		marshal, _ := json.Marshal(resp.Answer.BestAnswer)
		ans = string(marshal)
	}
	// 记录空答案或者有答案才会被记录
	if manager.GetManager().GetConfig().RecordEmptyAnswer || ans != "[]" || os.Getenv("SQL_DSN") != "" {
		t := entity.Tiku{
			Extra:    extra,
			Type:     int32(resp.Type),
			Question: resp.Question,
			Answer:   ans,
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
}
