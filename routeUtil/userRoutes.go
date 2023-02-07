package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/users/:uid", FetchUser())
	router.GET("/users", FetchUsers())
	router.DELETE("/users/:uid", DeleteUser())
	router.PUT("/users/:uid", EditUser())

}
