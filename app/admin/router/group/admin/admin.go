package admin

import (
	"Serve/app/admin/api"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(v *gin.RouterGroup) {
	v.GET("/get_admin_list", api.GetAdminList)
}
