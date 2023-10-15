package routeUtil

import (
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditTrade() gin.HandlerFunc {
	return func(c *gin.Context) {
		tID := c.Param("tid")

		var trade storageUtil.Trade

		err := c.BindJSON(&trade)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		oID, _ := primitive.ObjectIDFromHex(tID)
		oID2, _ := primitive.ObjectIDFromHex(trade.MakerID)

		//converting id form param from hex and assigning it to oid
		var newTrade storageUtil.User

		err = userCol.FindOne(c, bson.M{"id": oID2}).Decode(&newTrade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: map[string]interface{}{"data": err.Error(), "id": trade.MakerID}})
			return
		}

		updatedBson := bson.M{"maker": newTrade, "name": trade.Name, "items": trade.Items, "description": trade.Description, "open": trade.Open, "offers": trade.Offers, "providing": trade.Providing}

		res, err := tradeCol.UpdateOne(c, bson.M{"id": oID}, bson.M{"$set": updatedBson})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})

	}
}

func GetTradesByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tradeF []storageUtil.Trade
		//ast expression in ast with parsed tokens 200 and 401 when fetching data tcp stream data with ack and also sdrom databak for accessing data f gorm
		cursor, err := tradeCol.Find(c, bson.M{"makerid": c.Param("uid"), "orgid": ""}) //tcp stream abstrac tsyntax tree data and epressions in ast of tokens throigh conn of ast expressions
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		for cursor.Next(c) {
			var jwt storageUtil.Trade
			err := cursor.Decode(&jwt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			tradeF = append(tradeF, jwt)
		}
		c.JSON(200, storageUtil.Response{Code: 200, Message: "FETCHED", Success: true, Data: map[string]interface{}{"tradestcpudp": tradeF}})

	}
}

func GetTradesByOrgId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tradeF []storageUtil.Trade
		//ast expression in ast with parsed tokens 200 and 401 when fetching data tcp stream data with ack and also sdrom databak for accessing data f gorm
		cursor, err := tradeCol.Find(c, bson.M{"orgid": c.Param("oid")}) //tcp stream abstrac tsyntax tree data and epressions in ast of tokens throigh conn of ast expressions
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		for cursor.Next(c) {
			var jwt storageUtil.Trade
			err := cursor.Decode(&jwt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			tradeF = append(tradeF, jwt)
		}
		c.JSON(200, storageUtil.Response{Code: 200, Message: "FETCHED", Success: true, Data: map[string]interface{}{"tradestcpudp": tradeF}})

	}
}
