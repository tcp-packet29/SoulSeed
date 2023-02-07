package routeUtil

import (
	"errors"
	"fmt"
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

func LogAuth(username string, password string, c *gin.Context) (string, error) {
	var usr storageUtil.User
	err := userCol.FindOne(c, bson.M{"username": username}).Decode(&usr)
	if err != nil {
		return "", err
	}
	//didnt hash passwords so will use hashing check later
	if usr.Password != password {
		return "", errors.New("passwords do not match")
	}
	token, err := JWTGen(usr.Id.Hex())
	if err != nil {
		return "", err
	}
	return token, nil
}

func JWTGen(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix() //expriy time of hashmap
	fmt.Println(claims["exp"])
	claims["auth"] = true //time of creation
	claims["user_id"] = userId

	tokenString, err := token.SignedString(jwtEncryptionKey)
	if err != nil {
		return "", err
	}
	fmt.Println([]byte(tokenString))
	return tokenString, nil //no error

}

func VerifyJWTToken(c *gin.Context) (er error) { //takes in reuest handlign rfunc as param to add process and middlerwar

	if c.Request.Header["Token"] == nil {
		//if no jwt exists no auth
		c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized (no token exists)", Success: false, Data: nil})
		return errors.New("error in parsing jwt")
	} else {
		token, err := jwt.Parse(c.Request.Header["Token"][0], func(tok *jwt.Token) (interface{}, error) {
			_, valid := tok.Method.(*jwt.SigningMethodHMAC)
			if !valid {
				//if method is not dseigned method then no auth
				c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being from provider", Success: false, Data: nil})
				return nil, nil
			}
			return jwtEncryptionKey, nil
		})
		if err != nil {
			//if error in aprsing jwt no auth
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being parsed", Success: false, Data: map[string]interface{}{"error": err.Error()}})
			return errors.New("error in parsing jwt")
		}

		if token.Valid {
			return nil
		} else {
			//if invalid token no auth
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being valid due to some reason", Success: false, Data: nil})
			return errors.New("error in parsing jwt")
		}

	}
}

//TODO:
//1. Hashing and encoding passwords through bcrypt
//2. Getting fetcvhing current user from token and current and inuse active user by id from token
