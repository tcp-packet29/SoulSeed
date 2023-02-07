package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/genUtil"
	"main/middleware"
	auth "main/oAuthHandling"

	"main/routeUtil"
)

func main() {
	go middleware.RateLimitInit()
	router := gin.Default()

	router.Use(cors.Default()) //all,ow for creds in middleware request parsing and requersuit procedssign for ofdrs
	genUtil.ConnectToMongo()
	//mnot actually an instance of client we can actually reference, just for tersrting purposes

	access := router.Group("/access")
	access.POST("/login", auth.LoginHandle())
	access.POST("/users", routeUtil.CreateUser())
	authNeeded := router.Group("/app")
	authNeeded.Use(middleware.JwtAuth())
	routeUtil.UserRoutes(authNeeded)
	routeUtil.DealRoutes(authNeeded)
	routeUtil.OrgRoutes(authNeeded)
	//routes for different ufnctiaonlties
	router.Run(":8080")

}
