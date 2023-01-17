package oauthhandling

import (
	"context"
	"fmt"
	"main/genUtil"
	"main/storageUtil"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oAuthConfig = &oauth2.Config{
		ClientID: "",
		ClientSecret: "",
		RedirectURL: "http://localhost:8080/auth/google/callback",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	oAuthState = ""
)

func initOAUTH() {
	clientID, clientSecret := genUtil.GetGoogleData()
	oAuthConfig.ClientID = clientID
	oAuthConfig.ClientSecret = clientSecret
}