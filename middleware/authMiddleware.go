package middleware

import (
	"github.com/gin-gonic/gin"
	"main/routeUtil"
	"main/storageUtil"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := routeUtil.VerifyJWTToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, storageUtil.Response{Code: http.StatusUnauthorized, Message: "Unauthorized", Success: false, Data: map[string]interface{}{"data": err.Error()}})
			c.Abort() //abot thre request and requset processing
			return
		}
		c.Next() //continue the request

	}
}
