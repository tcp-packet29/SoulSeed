package routeUtil

import (
	"encoding/json"
	"main/genUtil"
	"main/storageUtil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtEncryptionKey = []byte(genUtil.GetJWTData())


func FetchUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ident := c.Param("uid")
		var userFound storageUtil.User

		oID, _ := primitive.ObjectIDFromHex(ident)
		//converting id form param from hex and assigning it to oid

		err := userCol.FindOne(c, bson.M{"id": oID}).Decode(&userFound) //finding user and decoding and transferring into userfound struct

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": userFound}})
	}
}

func JWTGen() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute) //expriy time of hashmap
	claims["iat"] = time.Now() //time of creation
	claims["user"] = "user"

	tokenString, err := token.SignedString()
}