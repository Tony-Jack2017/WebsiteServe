package service

import (
	"gorm.io/gorm"
)

type Service struct {
	Orm *gorm.DB
	Msg string
	MsgId string
	Error error
}