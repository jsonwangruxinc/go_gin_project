package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var token = "fas23dsfa5677fsa5523ffsfsfqwe=="

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("access_token")
		fmt.Println("access_token", accessToken)
		if accessToken != token {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "token 校验失败",
			})
			c.Abort()
		}
		c.Next()
	}
}
