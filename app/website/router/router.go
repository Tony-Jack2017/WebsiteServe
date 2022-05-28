package router

import (
	"Serve/app/website/router/group/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string) {
	gin.ForceConsoleColor()
	r := gin.Default()
	registerUserRouter(r)
	r.Run(":" + port)
}

func registerUserRouter(r *gin.Engine) {
	userGroup := r.Group("/apis/user")
	user.RegisterApisRouter(userGroup)
}