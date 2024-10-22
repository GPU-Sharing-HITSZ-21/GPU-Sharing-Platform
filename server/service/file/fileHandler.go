package file

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func HandleFileUpload(c *gin.Context) {
	// 获取请求头中的 JWT token 并解析出用户名
	token := c.Request.Header.Get("Authorization")
	username, err := utils.GetUsername(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: Invalid token"})
		return
	}

	log.Printf("用户名: %s\n", username)

	// 创建用户目录
	userDir := filepath.Join("/uploads/", username)

	// 打印用户目录日志
	log.Printf("用户目录: %s\n", userDir)

	// 检查目录是否存在
	if _, err := os.Stat(userDir); os.IsNotExist(err) {
		// 如果不存在，则创建目录
		if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建目录失败"})
			return
		}
	}

	// 读取上传的多个文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取文件失败"})
		return
	}

	files := form.File["files"] // 使用 "files" 作为表单字段名

	// 保存每个文件到用户目录
	for _, file := range files {
		if err := c.SaveUploadedFile(file, filepath.Join(userDir, file.Filename)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功"})
}
