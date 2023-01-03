package routeUtil

import (

	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)


func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ident := c.Param("uid")
		

		oID, _ := primitive.ObjectIDFromHex(ident)
		//converting id form param from hex and assigning it to oid

		res, err := userCol.DeleteOne(c, bson.M{"id": oID}) //finding user and decoding and transferring into userfound struct

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		if res.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})
	}
}