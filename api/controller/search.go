package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/internal/search"
	"net/http"
)

func Search(c *gin.Context) {
	var client = search.SearchClient{}

	var req model.SearchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	testRequest := model.SearchRequest{
		Question: "中国最美丽的风景",
	}

	// 调用被测试的方法
	response, err := client.Wanneng.SearchAnswer(testRequest)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}
