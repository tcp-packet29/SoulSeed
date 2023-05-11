package routeUtil

import (
	"context"
	ma "github.com/mailgun/mailgun-go/v4"
	"main/genUtil"
	"main/storageUtil"
	"time"

	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		orgId := c.Param("oid")
		var updatedOrg storageUtil.Organization

		err := c.BindJSON(&updatedOrg)
		if err != nil {
			fmt.Println("e")
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Body was not an organization", Success: false, Data: nil})
			//return reponse
			//send a response back to the client if error
			//return object of response struct
		}

		oID, err := primitive.ObjectIDFromHex(orgId)
		if err != nil {
			fmt.Println("e")
			//return response
		}
		fmt.Print(oID.Hex())
		fmt.Print(orgId)
		//will print the same thing

	}
}

func changeUserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//by ref not value also handelr ufnc to handle route
	}
}

func FetchOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//by ref not value also handelr ufnc to handle route
		id := c.Param("uid")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}
		var org storageUtil.Organization

		err = OrganizationCol.FindOne(c, bson.M{"id": oid}).Decode(&org)
		if err != nil {
			c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
			return
		}
		//could use genercis but getching with this code is fine
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"organization": org}})
	}
}

type orgVals struct {
	orgIds []string `bson:"orgIds"`
}

func MultOrg() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orgIdas storageUtil.OrgData
		err := c.BindJSON(&orgIdas)

		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		organizations := []storageUtil.Organization{}
		for _, val := range orgIdas.OrgVal { //jwt
			//jwt
			//jwt
			//jwt
			//jwt
			//jwt
			//nwt
			//jwt
			//jwt
			//jwt
			fmt.Println(val)
			oid, err := primitive.ObjectIDFromHex(val)
			if err != nil {
				c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"error": err.Error()}})
				return
			}
			var org storageUtil.Organization
			err = OrganizationCol.FindOne(c, bson.M{"id": oid}).Decode(&org)
			if err != nil {
				c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
				return
			}
			organizations = append(organizations, org)
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"organizations": organizations}})
	}
}

func AddUserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("oid")
		oid, err := primitive.ObjectIDFromHex(uid)
		var addedUser storageUtil.User
		err = c.BindJSON(&addedUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		organization := genUtil.FetchOrgById(uid, OrganizationCol, c, func() {
			c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})

		})

		fmt.Println(organization.Name)

		userList := organization.Users_ID
		userList = append(userList, addedUser.Id.Hex())

		finalOrganization := storageUtil.Organization{
			//new orgnaizaiton odel
			Id:          organization.Id,
			Name:        organization.Name,
			Description: organization.Description,
			Image:       organization.Image,
			Zipcode:     organization.Zipcode,
			Owner:       organization.Owner,
			Owner_ID:    organization.Owner_ID,
			Items:       organization.Items,
			Users:       append(organization.Users, addedUser), //or could jsut add user by idas9iofg or just add hrt eentire user
			Users_ID:    userList,
		}

		fmt.Println(finalOrganization.Users)

		//result, err := OrganizationCol.UpdateOne(c, bson.M{"id": uid}, bson.D{{"$set", bson.D{{"users", finalOrganization.Users}, {"users_id", finalOrganization.Users_ID}}}})
		result, err := OrganizationCol.ReplaceOne(c, bson.M{"id": oid}, finalOrganization)

		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": result}})

	}
}
func SendTokenMessage() gin.HandlerFunc { //strings instead of models for the sake of simplcity
	return func(c *gin.Context) {
		var model storageUtil.Model
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		org := genUtil.FetchOrgById(model.OrganizationCode, OrganizationCol, c, func() { c.JSON(404, gin.H{"error": "Organization not found"}) }) //redis key not found

		token := c.Param("token")
		aaa := c.Param("aaa")
		val := genUtil.GetMailgunData()

		mg := ma.NewMailgun("sandbox9c61bd1e4277496793c46e636d1f0a6c.mailgun.org", val)
		ca, cancel := context.WithTimeout(context.Background(), time.Second*14)
		defer cancel()
		if len(model.Emails) > 10 {
			c.JSON(400, gin.H{"error": "Too many emails"})
			return
		}
		for _, em := range model.Emails {
			tok := mg.NewMessage(
				"Gaurav <bansal22.gaurav@gmail.com>", //tobedecided
				"You have been invited to the organization "+org.Name,
				"",
				em,
			)

			tok.SetHtml("<h1>Please use the following token to join the organization: " + token + ".</h1><br>This token will expire in " + aaa + " hours")
			_, _, _ = mg.Send(ca, tok)
		}

	}

}

func UserInOrg() gin.HandlerFunc {
	return func(co *gin.Context) {
		oid, _ := primitive.ObjectIDFromHex(co.Param("oid"))

		var org storageUtil.Organization
		err := OrganizationCol.FindOne(co, bson.M{"id": oid}).Decode(&org)
		if err != nil {
			co.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
			return
		}
		in := false
		for _, id := range org.Users_ID {
			if id == co.Param("uid") {
				in = true
				break
			}
		}

		co.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"exists": in}})

	}
}

func AddItemOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("oid")
		var addedItem storageUtil.Item
		err := c.BindJSON(&addedItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		organization := genUtil.FetchOrgById(uid, OrganizationCol, c, func() {
			c.JSON(http.StatusNotFound, storageUtil.Response{Code: http.StatusNotFound, Message: "Not Found", Success: false, Data: nil})
		})
		//wt auth interpce thtp request
		//finding organization

		itemList := organization.Items
		itemList = append(itemList, addedItem) //no list of item ids since too cmplicated and items rae smaller just have lis of titme structs

		finalOrganization := storageUtil.Organization{
			Id:          organization.Id,
			Name:        organization.Name,
			Description: organization.Description,
			Image:       organization.Image,
			Zipcode:     organization.Zipcode,
			Owner:       organization.Owner,
			Owner_ID:    organization.Owner_ID,
			Items:       itemList,
			Users:       organization.Users, //or could jsut add user by idas9iofg
			Users_ID:    organization.Users_ID,
		}

		result, err := OrganizationCol.UpdateOne(c, bson.M{"id": uid}, bson.M{"$set": finalOrganization})
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
			return
		}

		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"data": result}})

	}
}

func Wotk() gin.HandlerFunc {
	return func(c *gin.Context) {
		x := c.Param("jwt")
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"jwt": x}})
	}
}

func startTimeout() {

}
