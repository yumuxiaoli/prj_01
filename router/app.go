package router

import (
	"prj_01/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", service.GetIndex)
	r.GET("/user", service.GetUserList)
	return r
}
