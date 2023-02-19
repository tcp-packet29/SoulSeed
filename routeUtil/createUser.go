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

var userCol *mongo.Collection = dbUtil.GetCollection("users")

func CreateUser() gin.HandlerFunc { //to be used in request handling suchj as POST
	return func(c *gin.Context) {
		var user storageUtil.User
		//since omitem,pty is used iod wikll nmot need to be specified leadign to it only being made here when cxreating user (easier for sending erq and making new user)
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//create user and add to dataabase aon post  func for creaitng userbased on body of request acreate and add ot datbase
		nUser := storageUtil.User{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			Items:    user.Items,
			Zipcode:  user.Zipcode,
		}

		resultIt, _ := userCol.Find(c, bson.M{}) //finding all of them, returning iterator
		//could do findone but then wouild check if mongo error and kind of iffy error conditional comparison
		defer resultIt.Close(c)

		for resultIt.Next(c) {
			var userFound storageUtil.User
			err := resultIt.Decode(&userFound)
			if err != nil {
				c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
				return
			}
			if userFound.Username == nUser.Username {
				c.JSON(http.StatusConflict, storageUtil.Response{Code: http.StatusConflict, Message: "Conflict", Success: false, Data: map[string]interface{}{"data": "Username already exists"}})
				return
			}
		} //slgithjly oinegfgiwecnt since leooping oiver all other users whenever new ueser but
		_, err := userCol.InsertOne(c, nUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusCreated, storageUtil.Response{Code: http.StatusCreated, Message: "Created", Success: true, Data: map[string]interface{}{"data": nUser}})

	}
}
