package main

import (
	"VOL/handler"
	"VOL/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 中间件
	router.Use(middleware.CORSMiddleware())

	// 定义路由和处理函数
	router.POST("/exec", handler.HandleCmd)
	router.POST("/k8s/command", handler.HandleK8sCommand)
	router.GET("/k8s/node", handler.HandlerGetNode)
	router.GET("/k8s/vcjob/status", handler.HandleVcjobStatus)
	router.GET("/k8s/pod_status", handler.HandlePodStatus)

	// 启动服务器
	router.Run(":8081")
}
