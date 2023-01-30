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

	tokenString, err := token.SignedString(jwtEncryptionKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil //no error
}

func verifyJWToken( handlFunc func(c *gin.Context)) gin.HandlerFunc { //takes in reuest handlign rfunc as param to add process and middlerwar
	return func (c *gin.Context) {
		if c.Header["Token"] == nil {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized (no token exists)", Success: false, Data: nil})
			return
		} else {
			token, err := jwt.Parse(c.Header["Token"][0], func(tok *jwt.Token) (interface{}, error) {
				_, valid :=token.Method.(*jwt.SigningMethodECDSA)
				if !valid {
					c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being from provider", Success: false, Data: nil})
					return nil, nil
				}
				return "", nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being parsed", Success: false, Data: nil})
				return
			}

			if token.Valid {
				handlFunc(c)
			} else {
				c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being valid due to some reason", Success: false, Data: nil})
				return
			}

			
		}
	}
}