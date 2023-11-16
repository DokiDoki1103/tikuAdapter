package controller

import (
	"code.sajari.com/docconv/v2"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/service"
	"github.com/itihey/tikuAdapter/pkg/global"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"net/http"
)

// Parser 解析文件接口 支持解析docx和xlsx
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
	contentType := file.Header.Get("Content-Type")
	if contentType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f, err := excelize.OpenReader(uploadedFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, global.ErrorParseFile)
			return
		}
		c.JSON(http.StatusOK, service.ParseXls(f))
	} else {
		convert, err := docconv.Convert(uploadedFile, contentType, true)
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
}
