package routeUtil

import (
	"encoding/base32"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/dbUtil"
	"main/storageUtil"
	"net/http"
	"crypto/rand"

)

var tokCol *mongo.Collection = dbUtil.GetCollection("orgCodes")


//creating id lsits

func CreateRandomId(c *gin.Context) (string, error) {
	var tokChecker storageUtil.Token
	byteL := make([]byte, 15)
	//create byte array
	//randomly read and ocnvert to stirng (random)
	_, err := rand.Read(byteL)
	if err != nil {
		return "", err
	}
	newVal := base32.StdEncoding.EncodeToString(byteL)
	obj := tokCol.FindOne(c, bson.M{"access": newVal}).Decode(&tokChecker)
	if obj != nil {
		return CreateRandomId(c)
	} else {
		return newVal, nil
	}


	//could just use mongodb id but its fine

}
func createAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		newTok := storageUtil.Token{
			ID: primitive.NewObjectID(),
			Access:           ,
			OrganizationCode: tok.OrganizationCode,
		}

		one, err := tokCol.InsertOne(c, tok)
		if err != nil {
			return
		}

		c.JSON(http.StatusCreated, storageUtil.Response{
			Code:    http.StatusCreated,
			Message: "Created obj",
			Success: false,
			Data: map[string]interface{}{
				"data": one,
			},
		})

	}
}

func FetchTok() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("uniqID")
	}
}
