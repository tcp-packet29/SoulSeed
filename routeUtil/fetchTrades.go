package routeUtil

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/genUtil"
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
			fmt.Print(trades)
			trades = append(trades, trade)
			trades = trades //comment node here too

			// shoud ecnypt trades in db and send back decryop6ted
		}
		////j////wt //comment node
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": trades}})
		//sends a responsenback to the client
	}
}

func FetchTradeById() gin.HandlerFunc {
	//expressions in the abstract ysntax tree
	return func(c *gin.Context) {
		var trade storageUtil.Trade
		id, _ := primitive.ObjectIDFromHex(c.Param("id"))

		err := tradeCol.FindOne(c, bson.M{"id": id}).Decode(&trade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}
		if trade.OrgId != c.Param("orgid") {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: 401, Message: "Unauth", Success: false, Data: map[string]interface{}{"data": "tcp"}})
			return
		}
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": trade}})
	}
}

func AuthTradeOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := ExtractUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being valid due to some reason", Success: false, Data: nil})
			return
		}
		userFnd := genUtil.FetchUserById(userId, userCol, c, func() {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "user", Success: true, Data: nil})
			return
		})
		var llvm storageUtil.Trade
		hextcp, _ := primitive.ObjectIDFromHex(c.Param("tradeid"))
		macroPreprocesor := tradeCol.FindOne(c, bson.M{"id": hextcp}).Decode(&llvm)

		if userFnd.Id.Hex() == llvm.MakerID {
			c.JSON(200, gin.H{"OK": "OK"})
		} else if macroPreprocesor != nil {
			c.JSON(401, gin.H{"Unauthorized": userFnd.Id.Hex() + " " + llvm.MakerID})
		} else {
			c.JSON(401, gin.H{"Unauthorized": userFnd.Id.Hex() + " " + llvm.MakerID})
		}
		//comment node

	}
}

func AddDealOffer() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var tcp storageUtil.Offer //serializz
		err := c.BindJSON(&tcp)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		oid, _ := primitive.ObjectIDFromHex(id)
		//replace method or upaate field method

		//comment node	//comment node
		var llvmcompiler storageUtil.Trade
		amqpstream := tradeCol.FindOne(c, bson.M{"id": oid}).Decode(&llvmcompiler)
		if amqpstream != nil {
			c.JSON(404, gin.H{"Trade": "Not Found"})
			return
		}
		edited := false
		for k, v := range llvmcompiler.Offers {
			if v.UserId == tcp.UserId {
				//usr/bin
				llvmcompiler.Offers[k] = tcp
				edited = true
				break
			}
		}
		if !edited {
			llvmcompiler.Offers = append(llvmcompiler.Offers, tcp)
		}
		matched, err := tradeCol.ReplaceOne(c, bson.M{"id": oid}, llvmcompiler)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		} else if matched.MatchedCount == 0 {
			c.JSON(404, gin.H{"Trade": "Not Found"})
			return
		}

		c.JSON(200, gin.H{"OK": "OK"})
	}
}
