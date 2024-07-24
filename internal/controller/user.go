package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/entity"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/util"
	"strconv"
)

// Auth 鉴权
func Auth(c *gin.Context) {
	claims, err := util.ParseJwtWithClaims(c.Request.Header.Get("Authorization"))

	if err != nil {
		logger.SysLog(err.Error())
		c.AbortWithStatusJSON(401, gin.H{
			"message": "请重新登录",
		})
		return
	}

	subject, err := claims.GetSubject()
	if err != nil {
		logger.SysLog(err.Error())
		c.AbortWithStatusJSON(401, gin.H{
			"message": "请重新登录",
		})
		return
	}

	if err != nil {
		logger.SysLog(err.Error())
		c.AbortWithStatusJSON(401, gin.H{
			"message": "请重新登录",
		})
		return
	}
	userID, _ := strconv.Atoi(subject)

	user, err := dao.User.Where(dao.User.ID.Eq(int32(userID))).First()
	if err != nil {
		logger.SysLog(err.Error())
		c.AbortWithStatusJSON(401, gin.H{
			"message": "请重新登录",
		})
		return
	}
	c.Set("user", user)
}

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	v, _ := c.Get(`user`)
	c.JSON(200, v.(*entity.User))
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	us := dao.User
	user, err := us.Where(us.Username.Eq(c.Query("username")), us.Password.Eq(c.Query("password"))).First()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}
	jwt := util.GenerateJwt(user.ID)

	if jwt == "" {
		c.JSON(400, gin.H{
			"message": "jwt生成失败",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "user login",
			"jwt":     jwt,
		})
	}
}

// UserList 自己开通的用户的列表
func UserList(c *gin.Context) {
	parent, _ := c.Get(`user`)
	users, err := dao.User.Where(dao.User.ParentID.Eq(parent.(*entity.User).ID)).Find()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "获取用户列表失败",
		})
		return
	}
	c.JSON(200, users)
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	parent, _ := c.Get(`user`)

	var user entity.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	// 检测用户名是否被使用
	_, err = dao.User.Where(dao.User.Username.Eq(user.Username)).First()
	if err == nil {
		c.JSON(400, gin.H{
			"message": "用户名已经被使用",
		})
		return
	}

	user.ParentID = parent.(*entity.User).ID

	err = dao.User.Create(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "创建用户失败",
		})
		return
	}
	c.JSON(200, user)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	parent, _ := c.Get(`user`)
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	user, err := dao.User.Where(dao.User.ID.Eq(int32(userID))).First()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "用户不存在",
		})
		return
	}

	if user.ParentID == parent.(*entity.User).ID {
		dao.User.Delete(user)
		c.JSON(200, gin.H{
			"message": "删除用户成功",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "无权限删除",
		})
	}
}
