package main

import (
	"github.com/gin-gonic/gin"
	"main/genUtil"
	"main/middleware"
	auth "main/oAuthHandling"
	"main/routeUtil"
)

func main() {
	go middleware.RateLimitInit()
	router := gin.Default()

	//all,ow for creds in middleware request parsing and requersuit procedssign for ofdrs
	genUtil.ConnectToMongo()
	router.Use(auth.NonContribCors())
	//mnot actually an instance of client we can actually reference, just for tersrting purposes

	access := router.Group("/access")
	tokens := router.Group("/tokens")
	access.Use(auth.NonContribCors())
	access.POST("/login", auth.LoginHandle())
	access.POST("/users", routeUtil.CreateUser())
	access.GET("/users/:username", routeUtil.FetchUserByUsername())
	access.GET("/users/token", routeUtil.FullID())
	routeUtil.TokRoutes(tokens)
	tokens.Use(routeUtil.TokenCheckMiddleware())
	authNeeded := router.Group("/app")
	authNeeded.Use(auth.NonContribCors())
	authNeeded.Use(middleware.JwtAuth())

	routeUtil.UserRoutes(authNeeded)
	routeUtil.DealRoutes(authNeeded)
	routeUtil.OrgRoutes(authNeeded)
	authNeeded.GET("/auth", auth.AuthEndpoint())
	//routes for different ufnctiaonlties
	router.Run(":8080")

}
