package router

import (
	"Serve/app/router/group/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string) {
	r := gin.Default()
	registerRouter(r)
	r.Run(":" + port)
}

func registerRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	user.RegisterUserRouter(userGroup)
}