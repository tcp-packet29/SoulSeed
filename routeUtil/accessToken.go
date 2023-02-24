package routeUtil

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/gin-gonic/gin"
	ma "github.com/mailgun/mailgun-go/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/dbUtil"
	"main/genUtil"
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

func TokenCheckMiddleware() gin.HandlerFunc { //better to check fi watnd token isi valid
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

		dur, _ := time.ParseDuration(num + "m") //llvm backend compiler

		newTok := storageUtil.Token{
			ID:               primitive.NewObjectID(),
			Access:           tokenToUse,
			OrganizationCode: tok.OrganizationCode,
			Expiry:           time.Now().Add(dur),
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
		if tokToFind.Expiry.Before(time.Now()) {
			_, err = tokCol.DeleteOne(c, bson.M{"id": tokToFind.ID})
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{
					Code:    500,
					Message: "internal server error",
					Success: false,
					Data: map[string]interface{}{
						"Error": err.Error(),
					},
				})
				return //ret
			}
			c.JSON(http.StatusNotAcceptable, storageUtil.Response{
				Code:    http.StatusNotAcceptable,
				Message: "Token expired",
				Success: false,
				Data: map[string]interface{}{
					"data": tokToFind,
				},
			})
			return //exit and escape handlerfunc for route
		}
		c.JSON(http.StatusOK, storageUtil.Response{
			Code:    http.StatusOK,  //emojis parsed o nfrontend
			Message: "Found object", //i love gcc
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

func SendConfirmationMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		em := c.Param("email")
		var usr storageUtil.User
		err := c.BindJSON(&usr)
		if err != nil {
			c.JSON(500, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}
		val, err := CreateRandomId(c)
		if err != nil {
			c.JSON(500, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}
		va := genUtil.GetMailgunData()
		mg := ma.NewMailgun("sandbox9c61bd1e4277496793c46e636d1f0a6c.mailgun.org", va)
		tok := mg.NewMessage(
			"Gaurav <bansal22.gaurav@gmail.com>",
			"Please confirm your email",
			"Hello "+usr.Username+",\n"+"Please use the following token to confirm your email: \n"+val+"\n\nIf you did not sign up for this service, please ignore this email.",
			em,
		)
		_, _, err = mg.Send(c, tok)
		if err != nil {
			c.JSON(500, storageUtil.Response{
				Code:    500,
				Message: "internal server error",
				Success: false,
				Data: map[string]interface{}{
					"Error": err.Error(),
				},
			})
		}
		c.JSON(200, storageUtil.Response{
			Code:    200,
			Message: "success",
			Success: true,
			Data: map[string]interface{}{
				"Token": val,
			},
		})
	}
}
