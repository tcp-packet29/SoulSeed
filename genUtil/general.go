package genUtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/storageUtil"
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
			Id:                userFound.Id,
			Username:          userFound.Username,
			Password:          "",
			Items:             userFound.Items,
			Zipcode:           userFound.Zipcode,
			OrganizationOwned: userFound.OrganizationOwned,
			OrganizationsIn:   userFound.OrganizationsIn,
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
		Id:                userFound.Id,
		Username:          userFound.Username,
		Email:             userFound.Email,
		Password:          "",
		Items:             userFound.Items,
		Zipcode:           userFound.Zipcode,
		OrganizationOwned: userFound.OrganizationOwned,
		OrganizationsIn:   userFound.OrganizationsIn,
	}

	return userCopy
}

func FetchOrgById(id string, col *mongo.Collection, ctx *gin.Context, f func(err error)) storageUtil.Organization {
	var oId primitive.ObjectID
	var userFound storageUtil.Organization
	var userCopy storageUtil.Organization

	oId, _ = primitive.ObjectIDFromHex(id)

	err := col.FindOne(ctx, bson.M{"id": oId}).Decode(&userFound) //finding user and decoding and transferring into userfound struct"}
	if err != nil {
		fmt.Println(err)
		f(err) //closure func tionm in a varioable to ivnvoke ervbent event
	}

	userCopy = storageUtil.Organization{
		Id:          userFound.Id,
		Name:        userFound.Name,
		Description: userFound.Description,
		Image:       userFound.Image,
		Zipcode:     userFound.Zipcode,
		Owner:       userFound.Owner,
		Owner_ID:    userFound.Owner_ID,
		Items:       userFound.Items,
		Users:       userFound.Users,
		Users_ID:    userFound.Users_ID,
	}

	return userCopy
	//i loe fetchinguser id and tookkt and wrapeprs111111111111111111111111111111111111111111111111111111 ytecode.
}
