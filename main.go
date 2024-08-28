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
	router.POST("/k8s/command", handler.HandleK8sCommand)
	router.GET("/k8s/node", handler.Handler_get_node)
	router.GET("/k8s/vcjob/status", handler.HandleVcjobStatus)
	router.GET("/k8s/pod_status", handler.HandlePodStatus) // 新增路由
	router.GET("/exec", handler.HandleCmd)

	// 启动服务器
	router.Run(":8081")
}
