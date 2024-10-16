package models

type User struct {
	Id       int
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"` // 注意：在生产环境中，请使用加密的密码
	Role     string `json:"role"`
}
