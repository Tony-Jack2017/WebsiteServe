package router

import (
	router "Serve/app/admin/router/group"
	router2 "Serve/app/website/router"
)

var adminInitRouter = router.InitRouter
var websiteInitRouter = router2.InitRouter

func InitRouter(appName string, port string)  {
	
	if appName == "admin" {
		adminInitRouter(port)
	}
	if appName == "website" {
		websiteInitRouter(port)
	}

}
