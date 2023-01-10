package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func DealRoutes(r *gin.Engine) {
	r.POST("/trades", CreateTrade())
	r.PUT("/trades/:tid", EditTrade())
	r.GET("/trades/:tid", FetchTrade())
	r.GET("/trades", FetchTrades())
}