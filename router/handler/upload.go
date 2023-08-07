package handler

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"path/filepath"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		responseData(c, err, nil)
		return
	}
	ext := filepath.Ext(file.Filename)
	const path = "/uploaded/"
	h := md5.New()
	tmpFile, err := file.Open()
	if err != nil {
		responseData(c, err, nil)
		return
	}
	fileBytes, err := io.ReadAll(tmpFile)
	if err != nil {
		responseData(c, err, nil)
		return
	}
	h.Write(fileBytes)
	name := hex.EncodeToString(h.Sum(nil))
	dst := "." + path + name + ext
	// 上传文件至指定的完整文件路径
	errSave := c.SaveUploadedFile(file, dst)
	log.Println("saveError:", errSave)
	responseData(c, errSave, map[string]string{"path": path, "filename": file.Filename})
}
