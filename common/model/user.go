package model

// BaseUser 使用者通用模型
type BaseUser struct {
	Account string `json:"account" gorm:"size:255;comment:账号"`
	Password string `json:"password" gorm:"size:255;comment:密码"`
	Salt string `json:"salt" gorm:"size:255;comment:加密所用的盐"`
	PasswordHash string `json:"password_hash" gorm:"size:255;comment:加密后的hash密码"`
}