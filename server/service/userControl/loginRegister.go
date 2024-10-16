package userControl

import (
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/dao/dataSource"
	"gpu-sharing-platform/models"
	"gpu-sharing-platform/utils"
	"net/http"
)

// RegisterUser 处理用户注册
func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置默认角色为普通用户
	newUser.Role = "USER"

	// 检查用户名是否已存在
	var existingUser models.User
	if err := dataSource.DB.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		return
	}

	// 在实际应用中，请添加密码加密
	dataSource.DB.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

// LoginUser 处理用户登录
func LoginUser(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := dataSource.DB.Where("username = ? AND password = ?", loginUser.Username, loginUser.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成令牌"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
