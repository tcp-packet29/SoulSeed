package routeUtil

import (
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchTrade() gin.HandlerFunc {
	return func(c *gin.Context) {
		tID := c.Param("tid")
		var trade storageUtil.Trade

		oID, _ := primitive.ObjectIDFromHex(tID)

		err := tradeCol.FindOne(c, bson.M{"id": oID}).Decode(&trade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": trade}})
	}
}