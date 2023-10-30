package routeUtil

import (
	"main/dbUtil"
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tradeCol *mongo.Collection = dbUtil.GetCollection("trades")
var DataCol = dbUtil.GetCollection("data")

func TCP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var datFound storageUtil.Data
		_ = DataCol.FindOne(c, bson.M{"id": "data"}).Decode(&datFound)
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": datFound.Map}})
	}
}

func CreateTrade() gin.HandlerFunc {
	return func(c *gin.Context) {
		var trade storageUtil.Trade

		err := c.BindJSON(&trade)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		var userFound storageUtil.User
		var organizationF storageUtil.Organization
		var datFound storageUtil.Data
		_ = DataCol.FindOne(c, bson.M{"id": "data"}).Decode(&datFound)
		oID, _ := primitive.ObjectIDFromHex(trade.MakerID)
		if trade.OrgId != "global" {
			oIdtcp, _ := primitive.ObjectIDFromHex(trade.OrgId)
			err = OrganizationCol.FindOne(c, bson.M{"id": oIdtcp}).Decode(&organizationF)
			if organizationF.Zipcode[:2] != c.Param("zipcode")[:2] {
				c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized", Success: false, Data: nil})
				return
			}
		}

		for _, val := range trade.Items {
			datFound.Map[val]++
		} //converting id form param from hex and assigning it to oid

		DataCol.ReplaceOne(c, bson.M{"id": "data"}, datFound)
		err = userCol.FindOne(c, bson.M{"id": oID}).Decode(&userFound) //finding user and decoding and transferring into userfound struct

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		newUser := storageUtil.User{
			Id:       userFound.Id,
			Username: userFound.Username,
			Email:    userFound.Email,
			Password: "",
			Items:    userFound.Items,
			Zipcode:  userFound.Zipcode,
		}

		//jmustg e3dit password on 8wserefound ifv fierld 9i8ns txcug tis not rewaDP0NLY BUT ALSO WRITE

		newTrade := storageUtil.Trade{
			Id:          primitive.NewObjectID(),
			Maker:       newUser,
			MakerID:     trade.MakerID,
			Name:        trade.Name,
			Items:       trade.Items,
			Description: trade.Description,
			Open:        true,
			OrgId:       trade.OrgId,
			Latlong:     trade.Latlong,
			Offers:      trade.Offers,
			Providing:   trade.Providing,
		}

		_, err = tradeCol.InsertOne(c, newTrade)

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		c.JSON(http.StatusCreated, storageUtil.Response{Code: http.StatusCreated, Message: "Created", Success: true, Data: map[string]interface{}{"data": newTrade}})
	}
}

func GetOfferAmount() gin.HandlerFunc {
	return func(c *gin.Context) {
		tID := c.Param("tid")

		var trade storageUtil.Trade
		dat, err := primitive.ObjectIDFromHex(tID)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: err.Error(), Success: false, Data: nil})
			return
		}
		err = tradeCol.FindOne(c, bson.M{"id": dat}).Decode(&trade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: err.Error(), Success: false, Data: nil})
			return
		} // couldnot find or error in fining

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "tcp", Success: true, Data: map[string]interface{}{"data": len(trade.Offers)}})

	}
}

func GetOffer() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		tid := c.Param("tid")

		var tradee storageUtil.Trade

		trade, _ := primitive.ObjectIDFromHex(tid)

		err := tradeCol.FindOne(c, bson.M{"id": trade}).Decode(&tradee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: err.Error(), Success: false, Data: nil})
			return
		}

		for _, val := range tradee.Offers {
			if val.UserId == id {
				c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "tcp", Success: true, Data: map[string]interface{}{"data": val}})
				return
			}
		}

		c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})

	}
}
