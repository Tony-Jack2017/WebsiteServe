package router

import (
	"Serve/app/admin/router/group/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	registerRouter(r)
	r.Run(":8080")
}

func registerRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	user.RegisterUserRouter(userGroup)
}