package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_project/web"
)

func InitCourse(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/course", web.GetCourse)
	v1.POST("/course", web.AddCourse)
}
