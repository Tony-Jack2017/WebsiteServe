package api

import "github.com/gin-gonic/gin"

// AdminLogin 管理员登录
// @Summary 管理员登录
// @Description 管理员登录，验证密码和账户名
// @Tags 登陆
// @Success 200 {object} response.Response{data=string,id=string,msg=string} "{"code": 200, "data": [...]}"
// @Router /api/apis_admin/admin/login [post]
func AdminLogin(c *gin.Context) {

}

func AddAdmin(c *gin.Context) {

}

func GetAdminList(c *gin.Context)  {
	
}
