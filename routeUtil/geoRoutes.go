package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func GeoRoutes(eng *gin.RouterGroup) {
	eng.GET("/geo/:zipcode", GettingId())
}
