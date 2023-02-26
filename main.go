package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dbUtil"
	"main/genUtil"
	"main/middleware"
	auth "main/oAuthHandling"
	"main/routeUtil"
	"main/storageUtil"
	"time"
)

func main() {
	_, err := genUtil.SendTokenMessage("sandbox9c61bd1e4277496793c46e636d1f0a6c.mailgun.org", storageUtil.Token{ID: primitive.NewObjectID(), Access: "1234", OrganizationCode: "1234", Expiry: time.Now().Add(time.Hour * 24)}, "Gaurav", "Org", "bansal22.gaurav@gmail.com")
	if err != "" {
		println(err)
	}
	go middleware.RateLimitInit()
	router := gin.Default()

	//all,ow for creds in middleware request parsing and requersuit procedssign for ofdrs
	dbUtil.ConnectToMongo()
	router.Use(auth.NonContribCors())
	//mnot actually an instance of client we can actually reference, just for tersrting purposes

	access := router.Group("/access")
	tokens := router.Group("/tokens")
	confirmation := router.Group("/email")
	access.Use(auth.NonContribCors())
	access.POST("/login", auth.LoginHandle())
	access.POST("/users", routeUtil.CreateUser())
	access.GET("/users/:username", routeUtil.FetchUserByUsername())
	access.GET("/users/token", routeUtil.FullID())
	routeUtil.TokRoutes(tokens)
	tokens.Use(routeUtil.TokenCheckMiddleware())
	confirmation.POST("/confirm/:email", routeUtil.SendConfirmationMessage()) //dont need middleware, just make random token
	authNeeded := router.Group("/app")
	authNeeded.Use(auth.NonContribCors())
	authNeeded.Use(middleware.JwtAuth())

	routeUtil.UserRoutes(authNeeded)
	routeUtil.DealRoutes(authNeeded)
	routeUtil.OrgRoutes(authNeeded)
	authNeeded.GET("/auth", auth.AuthEndpoint())
	//routes for different ufnctiaonlties
	//oi lve golang restapi guragavbansal restapis
	router.Run(":8080")

}
