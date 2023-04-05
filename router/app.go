package router

import (
	"prj_01/docs"
	"prj_01/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	idx := r.Group("/")
	{
		idx.GET("/index", service.GetIndex)
	}
	user := r.Group("/user")
	{
		user.GET("/", service.GetUserList)
		user.GET("/creatuser", service.CreatUser)
		user.GET("/deleteuser", service.DeleteUser)
		user.POST("/updateuser", service.UpdateUser)
	}
	return r
}
