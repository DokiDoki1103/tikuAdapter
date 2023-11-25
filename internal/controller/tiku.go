package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
)

type Page struct {
	PageNo   int            `json:"pageNo" form:"pageNo"`
	PageSize int            `json:"pageSize" form:"pageSize"`
	Total    int64          `json:"total" form:"total"`
	Items    []*entity.Tiku `json:"items" form:"items"`
}

func GetQuestions(c *gin.Context) {
	var page Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	fmt.Println(page)
	items, i, err := dao.Tiku.FindByPage(page.PageNo*page.PageSize, page.PageSize)
	if err != nil {
		return
	}
	page.Total = i
	page.Items = items
	c.JSON(200, page)
}
