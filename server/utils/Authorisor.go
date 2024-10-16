package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/dao/dataSource"
	"gpu-sharing-platform/models"
	"net/http"
)

func AuthorizeRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未提供令牌"})
			c.Abort()
			return
		}

		// 验证 token
		claims, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无效令牌"})
			c.Abort()
			return
		}

		// 从 token 中获取用户角色
		username := claims.Claims.(jwt.MapClaims)["username"].(string)
		var user models.User
		if err := dataSource.DB.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户未找到"})
			c.Abort()
			return
		}

		// 检查角色是否匹配
		for _, role := range roles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": "没有权限"})
		c.Abort()
	}
}
