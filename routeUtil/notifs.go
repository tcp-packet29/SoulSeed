package routeUtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dbUtil"
	"main/genUtil"
	"main/storageUtil"
	"net/http"
)

type notif struct {
	UserFor primitive.ObjectID `bson:"userFor,omitempty"`
	Body    string             `bson:"body"`
}

var notificationCol = dbUtil.GetCollection("notif")

func AddNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var notification notif
		err := c.ShouldBindJSON(&notification)
		if err != nil {
			c.JSON(400, gin.H{"400": "couldnt bind"})
		}
		_, err = notificationCol.InsertOne(c, notification)
		if err != nil {
			c.JSON(500, gin.H{"500": "couldnt insert"})
		}
		user := storageUtil.User{}
		hex, _ := primitive.ObjectIDFromHex(notification.UserFor.Hex())
		_ = userCol.FindOne(c, bson.M{"id": hex}).Decode(&user)
		fmt.Println(hex.Hex())

		err = genUtil.SendNotif("sandbox9c61bd1e4277496793c46e636d1f0a6c.mailgun.org", user.Username, notification.Body, user.Email)
		if err != nil {
			c.JSON(500, gin.H{"500": err.Error()})
		}
		c.JSON(200, gin.H{"codegen": "TCP Stream"})
	}
}

func GetNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := ExtractUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized due to jwt not being valid due to some reason", Success: false, Data: nil})
			return
		}
		userFnd := genUtil.FetchUserById(userId, userCol, c, func() {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "user", Success: true, Data: nil})
		})
		//automatically converst from hex tro oid

		var notifs []notif

		cursor, err := notificationCol.Find(c, bson.M{"userFor": userFnd.Id})
		if err != nil {
			c.JSON(500, gin.H{"500": "couldnt find"})
		}
		err = cursor.All(c, &notifs)
		if err == nil {
			c.JSON(200, notifs)
		} else {
			c.JSON(500, gin.H{"500": "couldnt find"})
		}
	}
}

func NotifRoutes(r *gin.RouterGroup) {
	r.GET("/notifs", GetNotifications())
	r.POST("/notifs/add", AddNotification())
}
