package routeUtil

import (
	"main/dbUtil"
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
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

		nUser := storageUtil.User{
			Id:       primitive.NewObjectID(),
			Username: user.Username,
			Password: user.Password,
		}

		_, err := userCol.InsertOne(c, nUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusCreated, storageUtil.Response{Code: http.StatusCreated, Message: "Created", Success: true, Data: map[string]interface{}{"data": nUser}})


	}
}