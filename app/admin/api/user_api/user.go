package user_api

import (
	"Serve/tools/color_print"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context)  {
	color_print.ColorPrintln(color_print.Success("test"))
}