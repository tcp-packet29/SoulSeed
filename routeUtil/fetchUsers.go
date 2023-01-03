package routeUtil

import (

	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	

)


func FetchUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var users []storageUtil.User

		

		resultIt, err := userCol.Find(c, bson.M{}) //finding all of them, returning iterator


		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		defer resultIt.Close(c)
		for resultIt.Next(c) {
			var userFound storageUtil.User
			err = resultIt.Decode(&userFound)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			users = append(users, userFound)
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": users}})
	}
}