package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/internal/middleware"
	"github.com/itihey/tikuAdapter/pkg/util"
	"strconv"
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
		tx = tx.Where(dao.Tiku.QuestionText.Like("%" + util.GetQuestionText(searchValue.Question) + "%"))
	}
	if searchValue.Extra != "" {
		tx = tx.Where(dao.Tiku.Extra.Like(searchValue.Extra))
	}
	if searchValue.Source != 0 {
		tx = tx.Where(dao.Tiku.Source.Eq(searchValue.Source))
	}

	if searchValue.OnlyShowEmptyAnswer {
		tx = tx.Where(dao.Tiku.Answer.Eq("[]"))
	}
	items, total, err := tx.FindByPage(searchValue.PageNo*searchValue.PageSize, searchValue.PageSize)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "服务器错误",
		})
		return
	}

	c.JSON(200, Page{
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
	updates, err := dao.Tiku.Where(dao.Tiku.ID.Eq(int32(id))).Updates(tiku)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "服务器错误",
		})
		return
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
		if tiku.Extra != "" {
			tiku.Source = 2
		} else {
			tiku.Source = 1
		}
		middleware.FillHash(tiku)
		err := dao.Tiku.Create(tiku)
		if err == nil {
			count++
		}
	}
	c.JSON(200, gin.H{
		"message": "成功创建" + strconv.Itoa(count) + "条数据",
	})
}
