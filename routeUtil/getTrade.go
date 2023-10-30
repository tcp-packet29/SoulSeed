package routeUtil

import (
	"main/genUtil"
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

func FetchTradeOffers() gin.HandlerFunc {
	return func(c *gin.Context) {
		tID := c.Param("tid")
		var trad storageUtil.Trade
		oId, _ := primitive.ObjectIDFromHex(tID)

		err := tradeCol.FindOne(c, bson.M{"id": oId}).Decode(&trad)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
		}
		val, err := ExtractUserID(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
		}
		if trad.MakerID != val {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: 401, Message: "Unauth", Success: false, Data: map[string]interface{}{"data": "jwt"}})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": trad.Offers}})

	}

}

func AcceptOffer() gin.HandlerFunc {
	return func(c *gin.Context) {
		tid := c.Param("tid")
		oid := c.Param("uid")
		message := storageUtil.Mess{}
		err := c.BindJSON(&message)
		if err != nil {
			c.JSON(500, storageUtil.Response{Code: 500, Message: err.Error(), Success: false, Data: nil})
			return
		}
		var trad storageUtil.Trade
		objident, _ := primitive.ObjectIDFromHex(tid)
		//err := tradeCol.ReplaceOne(c, bson.M{"id": objident}, bson.M{"$set": bson.M{"open": false, "accepted": true, "acceptedOffer": oid}})
		//get and then set
		err = tradeCol.FindOne(c, bson.M{"id": objident}).Decode(&trad)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error2", Success: false, Data: nil})
			return
		}
		val, err := ExtractUserID(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error1", Success: false, Data: nil})
			return
		}
		if trad.MakerID != val {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: 401, Message: "Unauth", Success: false, Data: map[string]interface{}{"data": "jwt"}})
			return
		}
		offer := storageUtil.Offer{}
		found := false
		for i, valt := range trad.Offers {
			if valt.UserId == oid {
				//discriminator
				trad.Offers[i].Approved = true
				found = true
				offer = valt
				break
			}
		}
		if !found {
			c.JSON(401, storageUtil.Response{Code: 401, Message: "Unauth", Success: false, Data: map[string]interface{}{"data": "jwt"}})
		}

		trad.Open = false
		_, err = tradeCol.ReplaceOne(c, bson.M{"id": objident}, trad)
		if err != nil {
			c.JSON(500, storageUtil.Response{Code: 500, Message: "Internal Server Error0", Success: false, Data: nil})
			return
		}
		err = genUtil.SendTcp("sandbox9c61bd1e4277496793c46e636d1f0a6c.mailgun.org", offer.Usr.Username, message.Message, trad.Maker.Email, offer.Usr.Email)
		msg := "OK"
		if err != nil {
			msg = err.Error()
		}
		c.JSON(200, storageUtil.Response{Code: 200, Message: msg, Success: true, Data: nil})

	}
}

func GetOfferItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		zipc := c.Param("zipc")

		trades := []storageUtil.Trade{}
		items := map[string]int{}
		iter, err := tradeCol.Find(c, bson.M{"maker.zipcode": zipc, "open": true, "providing": true})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		if err = iter.All(c, &trades); err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		for _, val := range trades {
			for _, valt := range val.Items {
				items[valt]++
			}
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": items}})
	}
}
