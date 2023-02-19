package routeUtil

import "github.com/gin-gonic/gin"

func TokRoutes(r *gin.RouterGroup) {
	r.POST("/tokens/:expiry", createAccessToken())
	r.GET("/tokens/:uniqID", GetToken())
}
