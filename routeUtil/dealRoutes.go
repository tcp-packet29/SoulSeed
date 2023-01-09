package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func DealRoutes(r *gin.Engine) {
	r.POST("/trades", CreateTrade())
}