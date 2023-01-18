package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func OrgRoutes(r *gin.Engine) {
	r.POST("/organizations", PostOrganization())
	r.PUT("/organizations/:oid/users", AddUserOrganization())
}