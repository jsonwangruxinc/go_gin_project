package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//根据模块划分接口
	//根据版本划分接口
	InitUser(r)
	InitCourse(r)
}
