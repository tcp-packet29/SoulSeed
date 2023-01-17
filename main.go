package main

import (
	
	"github.com/gin-gonic/gin"
	"main/genUtil"
	"main/routeUtil"

)

func main() {
	router := gin.Default()
	genUtil.ConnectToMongo()
	//mnot actually an instance of client we can actually reference, just for tersrting purposes
	routeUtil.UserRoutes(router)
	routeUtil.DealRoutes(router)
	router.Run(":8080")
}