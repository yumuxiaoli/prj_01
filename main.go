package main

import (
	"prj_01/router"
)

func main() {
	// 创建一个默认的路由引擎
	r := router.Router()

	// 配置路由

	// r.Run() 启动HTTP服务，默认在 0。0.0.0：8080 启动服务

	r.Run(":8080") //启动一个web服务
}
