package oauthhandling

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"main/genUtil"
	"main/routeUtil"
	"main/storageUtil"
	"net/http"
	"time"
)

type logInp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtKey = []byte(genUtil.GetJWTData())

type jwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func LoginHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inp logInp
		err := c.BindJSON(&inp)
		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		u := storageUtil.User{Username: inp.Username, Password: inp.Password}
		tok, err := routeUtil.LogAuth(u.Username, u.Password, c)

		if err != nil {
			c.JSON(http.StatusBadRequest, storageUtil.Response{Code: http.StatusBadRequest, Message: "Bad Request", Success: false, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Token": tok}) //sending token so can login
	}
}

func generateJWTWithClaims(username string) (tokens string, err error) {
	expir := time.Now().Add(1 * time.Hour)
	claims := &jwtClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expir.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	return tokenStr, err
}

func validToken(tok string) (err error) {
	toke, err := jwt.ParseWithClaims(tok, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return err
	}
	clai, ok := toke.Claims.(*jwtClaims)
	if !ok {
		err = errors.New("could not parse claims and params of jwt token")
		return err
	}
	if clai.ExpiresAt < time.Now().Unix() {
		err = errors.New("token/session expired")
		return err
	}
	return nil

}

func AuthEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "you are authenticated"})
	}
}

func NonContribCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
