package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/genUtil"
	"main/middleware"
	auth "main/oAuthHandling"
	"time"

	"main/routeUtil"
)

func main() {
	go middleware.RateLimitInit()
	router := gin.Default()

	router.Use(cors.Default()) //all,ow for creds in middleware request parsing and requersuit procedssign for ofdrs
	genUtil.ConnectToMongo()
	//mnot actually an instance of client we can actually reference, just for tersrting purposes

	access := router.Group("/access")
	access.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "token", "Token"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour}))
	access.POST("/login", auth.LoginHandle())
	access.POST("/users", routeUtil.CreateUser())
	access.GET("/users/:username", routeUtil.FetchUserByUsername())
	authNeeded := router.Group("/app")
	authNeeded.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "token", "Token"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour}))
	authNeeded.Use(middleware.JwtAuth())

	routeUtil.UserRoutes(authNeeded)
	routeUtil.DealRoutes(authNeeded)
	routeUtil.OrgRoutes(authNeeded)
	authNeeded.GET("/auth", auth.AuthEndpoint())
	//routes for different ufnctiaonlties
	router.Run(":8080")

}
