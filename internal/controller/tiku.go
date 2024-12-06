package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/internal/middleware"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"net/http"
	"strconv"
	"time"
)

// Page 分页
type Page struct {
	PageNo   int            `json:"pageNo" form:"pageNo"`
	PageSize int            `json:"pageSize" form:"pageSize"`
	Total    int64          `json:"total" form:"total"`
	Items    []*entity.Tiku `json:"items" form:"items"`
}

// SearchValue 搜索参数
type SearchValue struct {
	PageNo              int    `json:"pageNo" form:"pageNo"`
	PageSize            int    `json:"pageSize" form:"pageSize"`
	Source              int32  `json:"source" form:"source"`
	Extra               string `json:"extra" form:"extra"`
	OnlyShowEmptyAnswer bool   `json:"onlyShowEmptyAnswer" form:"onlyShowEmptyAnswer"`
	Question            string `json:"question" form:"question"`
	CourseName          string `json:"courseName" form:"courseName"`

	Type int32 `json:"type" form:"type"`
	Plat int32 `json:"plat" form:"plat"`
}

// GetQuestions 获取题库
func GetQuestions(c *gin.Context) {
	var searchValue SearchValue
	err := c.ShouldBindJSON(&searchValue)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	tx := dao.Tiku.Order(dao.Tiku.ID.Desc())
	if searchValue.Question != "" {
		tx.Where(dao.Tiku.Question.Like("%" + searchValue.Question + "%"))
	}
	if searchValue.Extra != "" {
		tx.Where(dao.Tiku.Extra.Like(searchValue.Extra))
	}

	if searchValue.CourseName != "" {
		tx.Where(dao.Tiku.CourseName.Eq(searchValue.CourseName))
	}

	if searchValue.Plat != -1 {
		tx.Where(dao.Tiku.Plat.Eq(searchValue.Plat))
	}

	if searchValue.Type != -1 {
		tx.Where(dao.Tiku.Type.Eq(searchValue.Type))
	}

	if searchValue.OnlyShowEmptyAnswer {
		tx.Where(dao.Tiku.Answer.Eq("[]"))
	}
	items, total, err := tx.FindByPage(searchValue.PageNo*searchValue.PageSize, searchValue.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
		return
	}

	c.JSON(http.StatusOK, Page{
		PageNo:   searchValue.PageNo,
		PageSize: searchValue.PageSize,
		Total:    total,
		Items:    items,
	})
}

// UpdateQuestions 更新题库
func UpdateQuestions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	var tiku *entity.Tiku
	err = c.ShouldBindJSON(&tiku)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	// 找到旧答案
	oldTiku, err := dao.Tiku.Where(dao.Tiku.ID.Eq(int32(id))).First()

	middleware.FillHash(tiku) // 更新题库的时候重新计算hash值
	updates, err := dao.Tiku.Where(dao.Tiku.ID.Eq(int32(id))).Updates(tiku)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "服务器错误",
		})
		return
	}
	if updates.RowsAffected == 1 {
		l := entity.Log{
			Qid:        oldTiku.ID,
			Action:     1,
			UserID:     0,
			CreateTime: time.Now(),
			OldAnswer:  oldTiku.Answer,
			NewAnswer:  tiku.Answer,
		}
		dao.Log.Create(&l)
	}
	c.JSON(200, gin.H{
		"data": updates,
	})
}

// DeleteQuestion 删除题目
func DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	// 记录删除日志
	oldTiku, err := dao.Tiku.Where(dao.Tiku.ID.Eq(int32(id))).First()
	marshal, _ := json.Marshal(oldTiku)
	l := entity.Log{
		Qid:        oldTiku.ID,
		Action:     2,
		UserID:     0,
		CreateTime: time.Now(),
		OldAnswer:  string(marshal),
	}
	dao.Log.Create(&l)

	dao.Tiku.Where(dao.Tiku.ID.Eq(int32(id))).Delete()
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// CreateQuestion 创建题目
func CreateQuestion(c *gin.Context) {
	var tikus []*entity.Tiku
	err := c.ShouldBindJSON(&tikus)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	var count = 0
	for _, tiku := range tikus {
		t := tiku
		middleware.FillHash(t)
		err := dao.Tiku.Create(t)
		if err == nil {
			marshal, _ := json.Marshal(t)
			l := entity.Log{
				Qid:        t.ID,
				Action:     0,
				UserID:     0,
				CreateTime: time.Now(),
				OldAnswer:  string(marshal),
			}
			err := dao.Log.Create(&l)
			if err != nil {
				logger.SysError(err.Error())
			}
			count++
		}
	}
	c.JSON(200, gin.H{
		"message": "成功创建" + strconv.Itoa(count) + "条数据",
	})
}

// Courses 获取这个平台的所有课程名称
func Courses(c *gin.Context) {
	tx := dao.Tiku.Select(dao.Tiku.CourseName)

	if c.Query("plat") != "" {
		i, err := strconv.Atoi(c.Query("plat"))
		if err != nil {
			c.String(400, "平台类型有错误")
			return
		}
		tx.Where(dao.Tiku.Plat.Eq(int32(i)))
	}
	find, err := tx.Group(dao.Tiku.CourseName).Find()
	if err != nil {
		return
	}
	courses := make([]string, 0)
	for i := range find {
		if find[i].CourseName != "" {
			courses = append(courses, find[i].CourseName)
		}
	}
	c.JSON(200, courses)
}
