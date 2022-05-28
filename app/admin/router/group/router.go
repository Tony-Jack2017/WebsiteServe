package router

import (
	"Serve/app/admin/router/group/admin"
	"Serve/common/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func InitRouter(port string) {
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.Default()
	r.Use(middleware.Cors())
	registerAdminRouter(r)
	r.Run(":" + port)
}

func registerAdminRouter(r *gin.Engine) {
	adminGroup := r.Group("/api/apis_admin")
	admin.RegisterAdminRouter(adminGroup)
}