package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("your_secret_key") // 请使用更强的密钥

// GenerateToken 生成 JWT
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // 72 小时过期

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken 验证 JWT
func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

// GetUsername 通过token获取username
func GetUsername(tokenStr string) (string, error) {
	token, err := ValidateToken(tokenStr)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "invalid token", err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "username not found in token claims", err
	}

	return username, nil
}
