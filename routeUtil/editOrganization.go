package routeUtil

import (

	"main/storageUtil"
	"main/genUtil"

	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"

	


)

func EditOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AddUserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("oid")
		var addedUser storageUtil.User
		err := c.BindJSON(&addedUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}
		organization := genUtil.FetchOrgById(uid, OrganizationCol, c, func() {c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
			
		})

		userList := organization.Users_ID
		userList = append(userList, addedUser.Id.Hex())

		finalOrganization := storageUtil.Organization{
			Id: organization.Id,
			Name: organization.Name,
			Description: organization.Description,
			Image: organization.Image,
			Zipcode: organization.Zipcode,
			Owner: organization.Owner,
			Owner_ID: organization.Owner_ID,
			Items: organization.Items,
			Users: append(organization.Users, addedUser), //or could jsut add user by idas9iofg or just add hrt eentire user
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
	return func(c *gin.Context) {
		uid := c.Param("oid")
		var addedItem storageUtil.Item
		err := c.BindJSON(&addedItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		organization := genUtil.FetchOrgById(uid, OrganizationCol, c, func() {c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})})
		//finding organization

		itemList := organization.Items
		itemList = append(itemList, addedItem) //no list of item ids since too cmplicated and items rae smaller just have lis of titme structs

		finalOrganization := storageUtil.Organization{
			Id: organization.Id,
			Name: organization.Name,
			Description: organization.Description,
			Image: organization.Image,
			Zipcode: organization.Zipcode,
			Owner: organization.Owner,
			Owner_ID: organization.Owner_ID,
			Items: itemList,
			Users: organization.Users, //or could jsut add user by idas9iofg
			Users_ID: organization.Users_ID,
		}

		result, err := OrganizationCol.UpdateOne(c, bson.M{"id": uid}, bson.M{"$set": finalOrganization})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": result}})



	}
}
