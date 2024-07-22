package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/registry/manager"
	"github.com/itihey/tikuAdapter/pkg/global"
	"io"
	"net/http"
)

// Plat -
func Plat(c *gin.Context) {
	c.JSON(http.StatusOK, manager.GetManager().GetConfig().Plat)
}

// UploadFile -
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorFileNotFound)
		return
	}
	parentDir := c.DefaultQuery("parentDir", "word")

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParseFile)
		return
	}
	defer src.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorFileHashError)
	}
	objectKey := fmt.Sprintf("%s/%s/%s", parentDir, hex.EncodeToString(hash.Sum(nil)), file.Filename)
	// 重新打开文件以保存
	src.Seek(0, io.SeekStart)
	err = manager.GetManager().GetBucket().PutObject(objectKey, src)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorFileUploadError)
		return
	}
	c.String(http.StatusOK, objectKey)
}
