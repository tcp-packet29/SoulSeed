package routeUtil

import (
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		specIdentifer := c.Param("uid")
		var user storageUtil.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		oID, _ := primitive.ObjectIDFromHex(specIdentifer)
		//converting id form param from hex and assigning it to oid
		resultIt, _ := userCol.Find(c, bson.M{}) //finding all of them, returning iterator

		for resultIt.Next(c) {
			var userFound storageUtil.User
			err := resultIt.Decode(&userFound)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			if userFound.Username == user.Username && specIdentifer != userFound.Id.Hex() {
				c.JSON(http.StatusConflict, storageUtil.Response{Code: http.StatusConflict, Message: "Username Already Exists; Conflict", Success: false, Data: map[string]interface{}{"data": specIdentifer, "other": userFound.Id.Hex()}})
				return
			}
		} //slgithjly oinegfgiwecnt since leooping oiver all other users whenever new ueser but fine

		updatedBSONRep := bson.M{"username": user.Username, "email": user.Email, "password": user.Password, "items": user.Items, "zipcode": user.Zipcode, "organization_owned": user.OrganizationOwned}

		res, err := userCol.UpdateOne(c, bson.M{"id": oID}, bson.M{"$set": updatedBSONRep})

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})
	}
}

func CreateOrg() gin.HandlerFunc {
	return func(c *gin.Context) {
		usrId := c.Param("ui")

		oId, _ := primitive.ObjectIDFromHex(usrId)

		var user storageUtil.User
		err := userCol.FindOne(c, bson.M{"id": oId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error, cannot decode", Success: false, Data: map[string]interface{}{"Data": 99999}})
			return
		}
		if len(user.OrganizationOwned) == 5 {
			c.JSON(http.StatusNotAcceptable, storageUtil.Response{Code: http.StatusOK, Message: "Organization Limit of 5 Exceeded", Success: false, Data: map[string]interface{}{"Data": 99999}})
			return
		}

		org, err := MakeOrg(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error, cannot decode", Success: false, Data: map[string]interface{}{"Data": 99999}})
			return
		}
		user.OrganizationOwned = append(user.OrganizationOwned, org.Id.Hex())
		c.JSON(http.StatusOK, Code: http.StatusInternalServerError, Message: "Internal Server Error, cannot decode", Success: false, Data: map[string]interface{}{"Data": 99999}}R)
	}
}

func AddOrgMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		usr := c.Param("uid")
		org := c.Param("rg")

		oID, _ := primitive.ObjectIDFromHex(usr)

		var user storageUtil.User

		err := userCol.FindOne(c, bson.M{"id": oID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error, cannot decode", Success: false, Data: map[string]interface{}{"hackclub": err.Error(), "reso": "ast"}})
			return
		}
		orgs := user.OrganizationsIn
		orgs = append(orgs, org)

		resalib, err := userCol.UpdateOne(c, bson.M{"id": oID}, bson.M{"$set": bson.M{"organizations_in": orgs}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error, cannot edit mdb doc", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: false, Data: map[string]interface{}{"Data": resalib}})

	}
}


