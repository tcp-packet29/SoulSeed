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
		ident := c.Param("uid")
		var user storageUtil.User
		

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		oID, _ := primitive.ObjectIDFromHex(ident)
		//converting id form param from hex and assigning it to oid
		resultIt, _ := userCol.Find(c, bson.M{}) //finding all of them, returning iterator

		for resultIt.Next(c) {
			var userFound storageUtil.User
			err := resultIt.Decode(&userFound)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			if userFound.Username == user.Username && ident != userFound.Id.Hex() {
				c.JSON(http.StatusConflict, storageUtil.Response{Code: http.StatusConflict, Message: "Username Already Exists; Conflict", Success: false, Data: map[string]interface{}{"data": ident, "other":userFound.Id.Hex()}})
				return
			}
		} //slgithjly oinegfgiwecnt since leooping oiver all other users whenever new ueser but fine

		updatedBSONRep := bson.M{"username": user.Username, "password": user.Password, "items": user.Items}

		res, err := userCol.UpdateOne(c, bson.M{"id" : oID}, bson.M{"$set": updatedBSONRep})

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})
	}
}