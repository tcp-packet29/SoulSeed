package routeUtil

import (
	"fmt"
	"main/dbUtil"
	"main/genUtil"
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrganizationCol *mongo.Collection = dbUtil.GetCollection("organizations")

func PostOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var organization storageUtil.Organization
		
		err := c.BindJSON(&organization) //binding from bodyu of re4quest
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		finalOrganization := storageUtil.Organization{
			Id: primitive.NewObjectID(),
			Name: organization.Name,
			Description: organization.Description,
			Image: organization.Image,
			Zipcode: organization.Zipcode,
			Owner: genUtil.FetchUserById(organization.Owner_ID, userCol, c, func() {fmt.Println("Error")}),
			Owner_ID: organization.Owner_ID,
			Items: organization.Items,
			Users: genUtil.FetchUsersByIDs(organization.Users_ID, userCol, c, func() {fmt.Println("Error")}), //or could jsut add user by idas9iofg
			Users_ID: organization.Users_ID,
		}
		
		_, err = OrganizationCol.InsertOne(c, finalOrganization)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": organization}})
	}
}