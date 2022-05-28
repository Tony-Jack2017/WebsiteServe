package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary 管理员相关操作
// @Description 管理员登录
// @Tags 登陆
// @Success 200 {object} response.Response{data=string,id=string,msg=string} "{"code": 200, "data": [...]}"
// @Router /api/apis_admin/admin/login [post]
func Login(c *gin.Context) {

}

// AddAdmin
// @Summary 管理员相关操作
// @Description 添加管理员
// @Tags 添加
// @Router /api/apis_admin/admin/add_admin [post]
func AddAdmin(c *gin.Context) {

}

// GetAdminList
// @Summary 管理员相关操作
// @Description 获取管理员信息列表
// @Tags 获取
// @Router /api/apis_admin/admin/get_admin_list [get]
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
func GetAdminList(c *gin.Context)  {
	resp := make(map[string]interface{}, 0)
	resp["admin_list"] = "hello world"
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取管理员列表", "data": resp})
}
