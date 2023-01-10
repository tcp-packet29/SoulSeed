package routeUtil

import (

	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	

)

func FetchTrades() gin.HandlerFunc {
	return func(c *gin.Context) {
		var trades []storageUtil.Trade

		iter, err := tradeCol.Find(c, bson.M{}) //finding all of them, returning iterator
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		for iter.Next(c) {
			var trade storageUtil.Trade
			err = iter.Decode(&trade)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			trades = append(trades, trade)
		}


		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": trades}})
	}
}