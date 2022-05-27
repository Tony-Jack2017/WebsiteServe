package model

import "Serve/common/model"

type Admin struct {
	AdminId int `json:"admin_id" gorm:"primaryKey;autoIncrement;comment:管理员Id"`
	AdminName string `json:"admin_name" gorm:"comment:管理员姓名"`
	Role string `json:"role" gorm:"comment:管理员角色"`
	Status string `json:"status" gorm:"comment:管理员状态"`
	Remark string `json:"remark" gorm:"comment:管理员备注"`
	Avatar string `json:"avatar" gorm:"comment:管理员头像"`
	Background string `json:"background" gorm:"comment:管理员背景封面"`
	model.BaseUser
	model.BaseModel
}
