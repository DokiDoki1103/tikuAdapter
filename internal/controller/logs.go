package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
)

// LogList - 日志列表
func LogList(c *gin.Context) {
	user, _ := c.Get("user")
	find, err := dao.Log.Where(dao.Log.UserID.Eq(user.(*entity.User).ID)).Where(dao.Log.Action.Neq(3)).Order(dao.Log.ID.Desc()).Limit(10).Find()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	c.JSON(200, find)
}
