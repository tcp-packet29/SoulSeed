package genUtil

import (

	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/storageUtil"
	"github.com/gin-gonic/gin"

)

func FetchUsersByIDs(ids []string, col *mongo.Collection, ctx *gin.Context, f func()) []storageUtil.User {
	var oId primitive.ObjectID
	var userFound storageUtil.User
	var userCopy storageUtil.User
	
	var users []storageUtil.User
	for _, id := range ids {
		oId, _ = primitive.ObjectIDFromHex(id)

		err := col.FindOne(ctx, bson.M{"id": oId}).Decode(&userFound) //finding user and decoding and transferring into userfound struct
		if err != nil {
			fmt.Println(err)
			f() //closure func tionm in a varioable to ivnvoke ervbent event
		}
		userCopy = storageUtil.User{
			Id: userFound.Id,
			Username: userFound.Username,
			Password: "",
			Items: userFound.Items,
			Zipcode: userFound.Zipcode,
		}

		}

		users = append(users, userCopy)

	return users
}

func FetchUserById(id string, col *mongo.Collection, ctx *gin.Context, f func()) storageUtil.User {
	var oId primitive.ObjectID
	var userFound storageUtil.User
	var userCopy storageUtil.User

	oId, _ = primitive.ObjectIDFromHex(id)

	err := col.FindOne(ctx, bson.M{"id": oId}).Decode(&userFound) //finding user and decoding and transferring into userfound struct"}
	if err != nil {
		fmt.Println(err)
		f() //closure func tionm in a varioable to ivnvoke ervbent event
	}

	userCopy = storageUtil.User{
		Id: userFound.Id,
		Username: userFound.Username,
		Password: "",
		Items: userFound.Items,
		Zipcode: userFound.Zipcode,
	}

	return userCopy
}

