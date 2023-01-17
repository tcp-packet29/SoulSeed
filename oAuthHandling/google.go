package oauthhandling

import (
	"context"
	"fmt"
	"main/genUtil"
	"main/storageUtil"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/url"
	"net/http"
)

var (
	oAuthConfig = &oauth2.Config{
		ClientID: "",
		ClientSecret: "",
		RedirectURL: "http://localhost:8080/oauth/google",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	oAuthState = ""
)

func googleLogin(c *gin.Context, oconf *oauth2.Config, oState string) {
	var user GoogleAuth
	URL, err := url.Parse(oconf.Endpoint.AuthURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, storageUtil.Response{Code: http.StatusInternalServerError, Message: "Internal Server Error", Success: false, Data: nil})
		return
	}
	

}

func initOAUTH() {
	clientID, clientSecret := genUtil.GetGoogleData()
	oAuthConfig.ClientID = clientID
	oAuthConfig.ClientSecret = clientSecret
}