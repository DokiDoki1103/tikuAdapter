package middleware

import (
	"encoding/json"
	"github.com/gookit/goutil/strutil"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/internal/registry/manager"
	"github.com/itihey/tikuAdapter/pkg/model"
	"sort"
	"strconv"
)

// FillHash 填充题库的hash值
func FillHash(t *entity.Tiku) {
	if t.Answer == "" {
		t.Answer = "[]"
	} else if t.Options == "" {
		t.Options = "[]"
	}

	options := make([]string, 0)
	err := json.Unmarshal([]byte(t.Options), &options)
	if err != nil {
		t.Options = "[]" // 如果解析失败，就设置为空数组
	}
	sort.Strings(options) // 将选项排序

	sortOptionsStr, err := json.Marshal(options)
	if err != nil {
		sortOptionsStr = []byte("[]")
	}

	t.Hash = strutil.Md5(t.Question + string(sortOptionsStr) + strconv.Itoa(int(t.Type)) + strconv.Itoa(int(t.Plat)))
}

// CollectAnswer 收集没有搜索到的答案
func CollectAnswer(resp model.SearchResponse, courseName, extra string) {
	sort.Strings(resp.Options) // 将选项排序
	opts, err := json.Marshal(resp.Options)
	if err != nil {
		opts = []byte("[]")
	}
	ans := "[]"
	if len(resp.Answer.AnswerKey) > 0 && len(resp.Answer.BestAnswer) > 0 { // 客观题能直接找到answerKey
		marshal, _ := json.Marshal(resp.Answer.BestAnswer)
		ans = string(marshal)
	} else if len(resp.Answer.BestAnswer) > 0 && resp.Type != 3 && resp.Type != 0 && resp.Type != 1 { // 排除客观题之后依然有答案
		marshal, _ := json.Marshal(resp.Answer.BestAnswer)
		ans = string(marshal)
	}
	// 记录空答案或者有答案才会被记录
	if manager.GetManager().GetConfig().RecordEmptyAnswer || ans != "[]" {
		t := entity.Tiku{
			CourseName: courseName,
			Extra:      extra,
			Type:       int32(resp.Type),
			Question:   resp.Question,
			Answer:     ans,
			Options:    string(opts),
			Plat:       int32(resp.Plat),
		}
		FillHash(&t)
		err := dao.Tiku.Create(&t)
		if err != nil {
			// 已经收录过，但是可能答案为空的情况，那么就需要去更新答案
			tk, errNotFind := dao.Tiku.Where(dao.Tiku.Hash.Eq(t.Hash)).First()
			if errNotFind != nil && tk.Answer == "[]" {
				dao.Tiku.Where(dao.Tiku.ID.Eq(tk.ID)).Update(dao.Tiku.Answer, ans)
			}
		}
	}
}
