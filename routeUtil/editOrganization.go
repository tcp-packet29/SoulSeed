package routeUtil

import (

	"main/storageUtil"
	"main/genUtil"

	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"

	"fmt"


)

func EditOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AddUserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("oid")

		organization := genUtil.FetchOrgById(uid, userCol, c, func() {c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
			return
		})

		userList := organization.Users_ID
		userList = append(userList, uid)

		finalOrganization := storageUtil.Organization{
			Id: organization.Id,
			Name: organization.Name,
			Description: organization.Description,
			Image: organization.Image,
			Zipcode: organization.Zipcode,
			Owner: organization.Owner,
			Owner_ID: organization.Owner_ID,
			Items: organization.Items,
			Users: genUtil.FetchUsersByIDs(userList, userCol, c, func() {fmt.Println("Error")}), //or could jsut add user by idas9iofg
			Users_ID: userList,
		}

		result, err := OrganizationCol.UpdateOne(c, bson.M{"id": uid}, bson.M{"$set": finalOrganization})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": result}})



	}
}

func AddItemOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
