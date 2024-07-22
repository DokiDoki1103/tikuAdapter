package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
)

// FileMd5 -
func FileMd5(f *multipart.FileHeader) (string, error) {
	src, err := f.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 创建一个哈希计算器
	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
