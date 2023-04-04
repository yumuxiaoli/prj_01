package router

import (
	"prj_01/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r1 := r.Group("/")
	{
		r1.GET("/index", service.GetIndex)
		r1.GET("/user", service.GetUserList)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
