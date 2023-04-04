package main

import (
	"prj_01/docs"
	"prj_01/router"
	"prj_01/utils"
)

// @title        RESTful API
// @version      1.0
// @description  RESTful API Document
// @host         127.0.0.1:8080/index

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	// 创建一个默认的路由引擎
	docs.SwaggerInfo.BasePath = ""
	r := router.Router()

	// 配置路由

	// r.Run() 启动HTTP服务，默认在 0。0.0.0：8080 启动服务

	r.Run(":8080") //启动一个web服务
}
