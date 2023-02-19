package routeUtil

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/dbUtil"
	"main/storageUtil"
	"net/http"
	"time"
)

var tokCol *mongo.Collection = dbUtil.GetCollection("orgCodes")

//could use ttl but i think on request is better

//creating id lsits

func CreateRandomId(c *gin.Context) (string, error) {
	var tokChecker storageUtil.Token
	byteL := make([]byte, 10)
	//create byte array
	//randomly read and ocnvert to stirng (random)
	_, err := rand.Read(byteL)
	if err != nil {
		return "", err
	}
	newVal := base32.StdEncoding.EncodeToString(byteL)[:10]

	_ = tokCol.FindOne(c, bson.M{"access": newVal}).Decode(&tokChecker)
	if tokChecker.Access == newVal {
		return CreateRandomId(c)
	} else {
		return newVal, nil
	}

	//could just use mongodb id but its fine

}

func TokenCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		checkCont, err := tokCol.Find(c, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}

		defer checkCont.Close(c)

		for checkCont.Next(c) {
			var tok storageUtil.Token
			err := checkCont.Decode(&tok)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{
					Code:    500,
					Message: "internal server error",
					Success: false,
					Data: map[string]interface{}{
						"Error": err.Error(),
					},
				})
			}

			if tok.Expiry.Before(time.Now()) {
				_, err := tokCol.DeleteOne(c, bson.M{"ID": tok.ID})
				if err != nil {
					c.JSON(http.StatusInternalServerError, storageUtil.Response{
						Code:    500,
						Message: "internal server error",
						Success: false,
						Data: map[string]interface{}{
							"Error": err.Error(),
						},
					})
				}
			}
		}

	}
}
func createAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		num := c.Param("expiry")
		var tok storageUtil.Token

		if err := c.BindJSON(&tok); err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}

		tokenToUse, err := CreateRandomId(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}

		dur, _ := time.ParseDuration(num + "m")

		newTok := storageUtil.Token{
			ID:               primitive.NewObjectID(),
			Access:           tokenToUse,
			OrganizationCode: tok.OrganizationCode,
			Expiry:           time.Now().Add(time.Minute * dur),
		}

		one, err := tokCol.InsertOne(c, newTok)
		if err != nil {
			return
		}

		c.JSON(http.StatusCreated, storageUtil.Response{
			Code:    http.StatusCreated,
			Message: "Created obj",
			Success: true,
			Data: map[string]interface{}{
				"data": one,
			},
		})

	}
}

func GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("uniqID")
		var tokToFind storageUtil.Token

		err := tokCol.FindOne(c, bson.M{"access": id}).Decode(&tokToFind)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}
		c.JSON(http.StatusOK, storageUtil.Response{
			Code:    http.StatusOK,
			Message: "Created obj",
			Success: true,
			Data: map[string]interface{}{
				"data": tokToFind,
			},
		})

	}
}

//func FetchTok() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id := c.Param("uniqID")
//	}
//}
