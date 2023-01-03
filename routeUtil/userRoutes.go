package routeUtil

import (
	
	
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users/:uid", FetchUser())
	router.GET("/users", FetchUsers())
	router.POST("/users", CreateUser())
	router.DELETE("/users/:uid", DeleteUser())
	router.PUT("/users/:uid", EditUser())
}