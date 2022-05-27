package model

import "Serve/common/model"

type User struct {
	UserId int `json:"user_id" gorm:"primaryKey;autoIncrement;comment:用户id"`
	UserName string `json:"user_name" gorm:"size:255;comment:用户名称"`
	Avatar string `json:"avatar" gorm:"size:255;comment:用户头像"`
	Background string `json:"background" gorm:"size:255;comment:用户背景封面"`
	Role string `json:"role" gorm:"size:255;comment:用户角色"`
	Remark string `json:"remark" gorm:"size:255;comment:用户备注"`
	Status string `json:"status" gorm:"size:255;comment:用户状态"`
	model.BaseUser
	model.BaseModel
}

func (User)TableName() string  {
	return "users"
}