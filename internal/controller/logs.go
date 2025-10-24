package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
)

// LogList - 日志列表
func LogList(c *gin.Context) {
	user, _ := c.Get("user")

	// 获取分页参数，默认值：page=1, pageSize=10
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 确保参数合法
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	query := dao.Log.Where(dao.Log.UserID.Eq(user.(*entity.User).ID)).Where(dao.Log.Action.Neq(3))

	// 查询总数
	total, err := query.Count()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "查询失败",
		})
		return
	}

	// 查询分页数据
	find, err := query.Order(dao.Log.ID.Desc()).Limit(pageSize).Offset(offset).Find()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"data":     find,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
