package model

import "time"

type BaseModel struct {
	CreatedAt time.Time  `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"comment:修改时间"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"comment:删除时间"`
}

type Page struct {
	Total int `json:"total"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}
