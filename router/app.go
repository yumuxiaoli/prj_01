package router

import (
	"prj_01/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
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
		user.POST("/findandwd", service.FindUserByNameAndPwd)

		user.GET("/SendMsg", service.SendMsg)
	}
	return r
}
