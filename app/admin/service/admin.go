package service

import (
	"Serve/app/admin/model"
	model2 "Serve/common/model"
	"Serve/common/service"
	"errors"
)

type AdminService struct {
	service.Service
}

// GetAdminBy 获取管理员
func (s *AdminService) GetAdminBy() error {
	return nil
}

// GetAdminListBy 获取管理员列表
func (s *AdminService) GetAdminListBy() error {
	return nil
}

// AddAdmin 添加管理员
func (s *AdminService) AddAdmin(admin model.Admin) model2.BaseError {
	var err model2.BaseError
	var count int64
	err.Reason = s.Orm.Model(&admin).Where("username = ?", admin.AdminName).Count(&count)
	if err.Reason != nil {
		err.Describe = "查询用户出错"
		return err
	}
	if count > 0 {
		err.Describe = "创建用户失败"
		err.Reason = errors.New("用户已存在")
		return err
	}
	err.Reason = s.Orm.Create(&admin)
	if err.Reason != nil {
		return err
	}
}