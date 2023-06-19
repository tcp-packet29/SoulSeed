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

		oID, _ := primitive.ObjectIDFromHex(trade.MakerID)
		if trade.OrgId != "global" {
			oIdtcp, _ := primitive.ObjectIDFromHex(trade.OrgId)
			err = OrganizationCol.FindOne(c, bson.M{"id": oIdtcp}).Decode(&organizationF)
			if organizationF.Zipcode[:2] != c.Param("zipcode")[:2] {
				c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized", Success: false, Data: nil})
				return
			}
		}

		//converting id form param from hex and assigning it to oid

		err = userCol.FindOne(c, bson.M{"id": oID}).Decode(&userFound) //finding user and decoding and transferring into userfound struct

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		newUser := storageUtil.User{
			Id:       userFound.Id,
			Username: userFound.Username,
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
		}

		_, err = tradeCol.InsertOne(c, newTrade)

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		c.JSON(http.StatusCreated, storageUtil.Response{Code: http.StatusCreated, Message: "Created", Success: true, Data: map[string]interface{}{"data": newTrade}})
	}
}
