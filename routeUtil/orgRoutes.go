package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func OrgRoutes(r *gin.RouterGroup) {
	r.POST("/organizations", PostOrganization())
	r.PUT("/organizations/:oid/users", AddUserOrganization())
	r.GET("/organizations/:uid", FetchOrganization())
	r.PUT("/organizations/:oid/items", AddItemOrganization()) //put juts for okgging purposes, cna do soemyhing else ufncitoanltiy routes loginging peurposes weh logging debug mode and orgnaization log
	r.GET("/organizations/aa/:jwt", Wotk())
	r.POST("/organizations/idBulk", MultOrg())
	r.GET("/organizations/stuff/:oid/:uid", UserInOrg())
	r.GET("/organizations/admin/:oid/:uid", UserOwner())
	r.GET("/organizations/loc/:uid/:zipcode", CheckOrgLocation())
	r.POST("/organizations/sendRequest/:oid", SendRequestToJoin())

}
