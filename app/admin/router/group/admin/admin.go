package admin

import (
	"Serve/app/admin/api/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(v *gin.RouterGroup) {
	v.POST("/admin/login", admin.Login)
	v.POST("/admin/add_admin", admin.AddAdmin)
	v.GET("/admin/get_admin_list", admin.GetAdminList)
}
