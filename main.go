package main

import (
	"github.com/gin-gonic/gin"
	"main/dbUtil"
	"main/middleware"
	auth "main/oAuthHandling"
	"main/routeUtil"
)

func main() {

	dbUtil.UploadImage("./assetsFS/test.png", "test.png")
	go middleware.RateLimitInit()
	router := gin.Default()
	//create new router to be used with the backend webserver 	This involves creating a new handler for routes and group of routes gaurav bansal jwt auth middleware is on the use l low level systems uilzie proceser through assevmhly (7)
	//all,ow for creds in middleware request parsing and requersuit procedssign for ofdrs
	dbUtil.ConnectToMongo()
	dbUtil.ConnectToRedis()
	router.Use(auth.NonContribCors())
	//mnot actually an instance of client we can actually reference, just for tersrting purposes

	access := router.Group("/access")
	tokens := router.Group("/tokens")
	confirmation := router.Group("/email")
	confirmation.POST("/join/:org/:receiver/:token/:aaa", routeUtil.SendTokenMessage())
	access.Use(auth.NonContribCors())
	access.POST("/login", auth.LoginHandle())
	access.POST("/users", routeUtil.CreateUser())
	access.GET("/users/:username", routeUtil.FetchUserByUsername())
	access.GET("/users/token", routeUtil.FullID())
	access.GET("/users/organizations/:uid/:rg", routeUtil.AddOrgMember())
	routeUtil.TokRoutes(tokens)
	tokens.Use(routeUtil.TokenCheckMiddleware())
	confirmation.POST("/confirm/:email", routeUtil.SendConfirmationMessage()) //dont need middleware, just make random token
	confirmation.POST("/password/:email", routeUtil.SendPasswordMessage())
	authNeeded := router.Group("/app")
	authNeeded.Use(auth.NonContribCors())
	authNeeded.Use(middleware.JwtAuth())

	routeUtil.GeoRoutes(access)

	routeUtil.UserRoutes(authNeeded)
	routeUtil.DealRoutes(authNeeded)
	routeUtil.OrgRoutes(authNeeded)

	authNeeded.GET("/auth", auth.AuthEndpoint())
	//routes for different ufnctiaonlties
	//oi lve golang restapi guragavbansal restapis
	router.Run(":8080")
	//ythis is the abstract syntrax tree is a comment node

}
