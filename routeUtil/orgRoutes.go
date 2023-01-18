package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func OrgRoutes(r *gin.Engine) {
	r.POST("/organizations", PostOrganization())
	r.PUT("/organizations/:oid/users", AddUserOrganization())
	r.PUT("/organization/:oid/items", AddItemOrganization()) //put juts for okgging purposes, cna do soemyhing else ufncitoanltiy routes loginging peurposes weh logging debug mode and orgnaization log
} 