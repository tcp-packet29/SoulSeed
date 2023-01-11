package routeUtil

import (
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditTrade() gin.HandlerFunc {
	return func(c * gin.Context) {
		tID := c.Param("tid")

		var trade storageUtil.Trade
		
		err := c.BindJSON(&trade)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		oID, _ := primitive.ObjectIDFromHex(tID)
		var userFound storageUtil.User


		
		//converting id form param from hex and assigning it to oid

		err = userCol.FindOne(c, bson.M{"id": oID}).Decode(&userFound) //finding user and decoding and transferring into userfound struct

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		newUser := storageUtil.User{
			Id: userFound.Id,
			Username: userFound.Username,
			Password: "",
			Items: userFound.Items,
			Zipcode: userFound.Zipcode,
		}


		updatedBson := bson.M{"maker": newUser, "name": trade.Name, "items": trade.Items, "description": trade.Description, "open": trade.Open}

		res, err := tradeCol.UpdateOne(c, bson.M{"id": oID}, bson.M{"$set": updatedBson})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})

	}
}