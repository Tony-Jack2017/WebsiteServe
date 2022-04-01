package user

import (
	"Serve/app/admin/api/user_api"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(admin *gin.RouterGroup) {
	admin.GET("/login", user_api.GetUserList)
}
