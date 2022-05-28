package user

import (
	"github.com/gin-gonic/gin"
)

// SignIn
// @Description 用户注册
// @Summary 用户相关操作
// @Tags 添加
// @Router /api/apis/user/sign_in
func SignIn(c *gin.Context) {

}

// Login
// @Description 用户登录
// @Summary 用户相关操作
// @Tags 登录
// @Router /api/apis/user/login
func Login(c *gin.Context)  {

}

// GetUserList
// @Description 获取用户列表
// @Summary 用户相关操作
// @Tags 获取
// @Param username query string false "username"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/apis/user/get_user_list [get]
func GetUserList(c *gin.Context)  {

}