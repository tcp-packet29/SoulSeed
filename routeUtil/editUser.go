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

		updatedBSONRep := bson.M{"username": user.Username, "password": user.Password}

		res, err := userCol.UpdateOne(c, bson.M{"id" : oID}, bson.M{"$set": updatedBSONRep})

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})
	}
}