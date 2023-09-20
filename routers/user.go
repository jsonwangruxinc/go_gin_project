package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_project/Middleware"
	"go_gin_project/web"
)

func InitUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(Middleware.Auth())
	v1.GET("/user", web.GetUser)
	v1.POST("/user", web.AddUser)

}
