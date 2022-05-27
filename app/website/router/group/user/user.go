package user

import (
	"Serve/app/website/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterApisRouter(admin *gin.RouterGroup) {
	admin.GET("/login", user.Login)
}
