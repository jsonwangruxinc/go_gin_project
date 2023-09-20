package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(context *gin.Context) {
	id := context.DefaultQuery("key", "10")
	context.JSON(http.StatusOK, gin.H{"id": id})
}

type Course struct {
	Name    string  `form:"name" binding:"required"`
	Teacher string  `form:"teacher" binding:"required"`
	Price   float64 `form:"price" binding:"required"`
}

func AddCourse(contxet *gin.Context) {
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
