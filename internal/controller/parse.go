package controller

import (
	"code.sajari.com/docconv/v2"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/pkg/global"
	"mime/multipart"
	"net/http"
)

// Parse 解析文件接口
func Parser(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParam)
		return

	}
	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParam)
		return
	}
	defer func(uploadedFile multipart.File) {
		err := uploadedFile.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, global.ErrorParam)
			return
		}
	}(uploadedFile)

	convert, err := docconv.Convert(uploadedFile, file.Header.Get("Content-Type"), true)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParseFile)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, global.ErrorParseFile)
		return
	}

	c.JSON(http.StatusOK, convert)
}
