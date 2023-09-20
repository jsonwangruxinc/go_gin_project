package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(context *gin.Context) {
	id := context.DefaultQuery("key", "0")
	context.JSON(http.StatusOK, gin.H{"id": id})
}

type User struct {
	name  string `form:"name" binding:"required"`
	phone string `form:"phone" binding:"required,e164"`
}

func AddUser(contxet *gin.Context) {
	req := &Course{}
	err := contxet.ShouldBind(req)
	if err != nil {
		contxet.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	contxet.JSON(http.StatusOK, req)
}
