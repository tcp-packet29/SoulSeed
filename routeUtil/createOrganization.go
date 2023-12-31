package routeUtil

import (
	"fmt"
	"main/dbUtil"
	"main/genUtil"
	"main/storageUtil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrganizationCol *mongo.Collection = dbUtil.GetCollection("organizations")

func PostOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var organization storageUtil.Organization

		err := c.BindJSON(&organization) //binding from bodyu of re4quest
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: nil})
			return
		}

		finalOrganization := storageUtil.Organization{
			Id:          primitive.NewObjectID(), //newobj id when creatin new organization (new omngo objectif to referennce)cd
			Name:        organization.Name,
			Description: organization.Description,
			Image:       organization.Image,
			Zipcode:     organization.Zipcode,
			Owner:       genUtil.FetchUserById(organization.Owner_ID, userCol, c, func() { fmt.Println("Error") }),
			Owner_ID:    organization.Owner_ID,
			Items:       organization.Items,
			//macro actions stored in fif oqueue
			Users:    genUtil.FetchUsersByIDs(organization.Users_ID, userCol, c, func() { fmt.Println("Error") }), //or could jsut add user by idas9iofg
			Users_ID: organization.Users_ID,
		}

		_, err = OrganizationCol.InsertOne(c, finalOrganization)
		if err != nil {
			c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: map[string]interface{}{
				"data": err, //shows err imnr espo0nse to  eb read
				"info": finalOrganization},
			//organization data fetched fro nomrganization, erortr in interl aerror wheninsertion into db
			})
			return
		}
		c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"orgFetched": organization}})
	}
}

func MakeOrg(c *gin.Context) (storageUtil.Organization, error) {

	var organization storageUtil.Organization

	err := c.BindJSON(&organization) //binding from bodyu of re4quest
	if err != nil {
		return storageUtil.Organization{}, err
	}

	finalOrganization := storageUtil.Organization{
		Id:          primitive.NewObjectID(), //newobj id when creatin new organization (new omngo objectif to referennce)cd
		Name:        organization.Name,
		Description: organization.Description,
		Image:       organization.Image,
		Zipcode:     genUtil.FetchUserById(organization.Owner_ID, userCol, c, func() { fmt.Println("Error") }).Zipcode,
		Owner:       genUtil.FetchUserById(organization.Owner_ID, userCol, c, func() { fmt.Println("Error") }),
		Owner_ID:    organization.Owner_ID,
		Items:       organization.Items,
		//macro actions stored in fif oqueue
		Users:    genUtil.FetchUsersByIDs(organization.Users_ID, userCol, c, func() { fmt.Println("Error") }), //or could jsut add user by idas9iofg
		Users_ID: organization.Users_ID,
	}

	_, err = OrganizationCol.InsertOne(c, finalOrganization)
	if err != nil {
		return finalOrganization, err
	}
	//c.JSON(http.StatusOK, storageUtil.Response{Code: http.StatusOK, Message: "OK", Success: true, Data: map[string]interface{}{"orgFetched": finalOrganization}})
	return finalOrganization, nil //99999
}
