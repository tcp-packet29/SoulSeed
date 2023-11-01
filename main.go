package main

import (
	"github.com/gin-gonic/gin"
	"main/dbUtil"
	"main/middleware"
	auth "main/oAuthHandling"
	"main/routeUtil"
	"net/http"
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

	//routeUtil.DataCol.InsertOne(context.Background(), bson.M{"id": "data", "map": map[string]int{"test": 1}})
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
	access.GET("/tradesin/:zipc", routeUtil.GetOfferItems())
	confirmation.POST("/confirm/:email", routeUtil.SendConfirmationMessage()) //dont need middleware, just make random token
	confirmation.POST("/password/:email", routeUtil.SendPasswordMessage())
	authNeeded := router.Group("/app")
	http.HandleFunc("/wc", routeUtil.Wsock())
	authNeeded.Use(auth.NonContribCors())
	access.GET("/allstuff", routeUtil.TCP())
	access.GET("/organizations/organizationsInArea/:zcode", routeUtil.GetOrganizationByZipcode())
	authNeeded.Use(middleware.JwtAuth())

	routeUtil.GeoRoutes(access)

	routeUtil.UserRoutes(authNeeded)
	//comments in abstrac syuntax tree
	routeUtil.DealRoutes(authNeeded)
	routeUtil.NotifRoutes(authNeeded)
	routeUtil.OrgRoutes(authNeeded)

	authNeeded.GET("/auth", auth.AuthEndpoint())
	authNeeded.POST("/jwt", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"jwtauthwebsockettcp": "webrtcrealrtiemstream"})
	})
	//routes for different ufnctiaonlties
	//oi lve golang restapi guragavbansal restapis
	router.Run(":8080")
	//ythis is the abstract syntrax tree is a comment node

}
