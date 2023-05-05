package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	ginServer := gin.Default()
	// 配置路由
	ginServer.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "Hello world!",
		})
	})
    // 启动 HTTP 服务,默认在 0.0.0.0:8080启动服务
	ginServer.Run()
    // ginServer.Run(":8081")	使用8081端口启动服务
}
