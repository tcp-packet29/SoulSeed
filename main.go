package main

import (
	
	"github.com/gin-gonic/gin"
	"main/genUtil"
	"main/routeUtil"
	"main/middleware"

)

func main() {
	go middleware.RateLimitInit()
	router := gin.Default()
	genUtil.ConnectToMongo()
	//mnot actually an instance of client we can actually reference, just for tersrting purposes
	routeUtil.UserRoutes(router)
	routeUtil.DealRoutes(router)
	routeUtil.OrgRoutes(router)
	router.Run(":8080")
}