package routeUtil

import (
	
	
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user/:uid", FetchUser())
	router.GET("/users", FetchUsers())
	router.POST("/user", CreateUser())
	router.DELETE("/user/:uid", DeleteUser())
	router.PUT("/user/:uid", EditUser())
}