package main

import (
	
	"github.com/gin-gonic/gin"
	"main/genUtil"
	"main/routeUtil"
	"main/middleware"
	"github.com/gin-contrib/cors"

)

func main() {
	go middleware.RateLimitInit()
	router := gin.Default()

	router.Use(cors.Default())//all,ow for creds in middleware request parsing and requersuit procedssign for ofdrs
	genUtil.ConnectToMongo()
	//mnot actually an instance of client we can actually reference, just for tersrting purposes
	routeUtil.UserRoutes(router)
	routeUtil.DealRoutes(router)
	routeUtil.OrgRoutes(router)
	//routes for different ufnctiaonlties
	router.Run(":8080")
}