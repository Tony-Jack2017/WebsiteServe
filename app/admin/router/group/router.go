package router

import (
	"Serve/app/admin/router/group/admin"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string) {
	r := gin.Default()
	registerAdminRouter(r)
	r.Run(":" + port)
}

func registerAdminRouter(r *gin.Engine) {
	adminGroup := r.Group("/apis_admin")
	admin.RegisterAdminRouter(adminGroup)
}