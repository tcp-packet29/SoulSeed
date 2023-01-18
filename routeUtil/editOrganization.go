package routeUtil

import (
	"main/dbUtil"
	"main/storageUtil"
	"main/genUtil"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"


)

func EditOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AddUserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("oid")

		user := genUtil.FetchUserById(uid, userCol, c, func() {c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
			return
		})

		

	}
}

func AddItemOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
