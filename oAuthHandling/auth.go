package oauthhandling

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"main/genUtil"
	"time"
)

var jwtKey = []byte(genUtil.GetJWTData())

type jwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
		err = errors.New("Could not parse claims and params of jwt token")
		return err
	}
	if clai.ExpiresAt < time.Now().Unix() {
		err = errors.New("Token/session expired")
		return err
	}
	return nil

}
