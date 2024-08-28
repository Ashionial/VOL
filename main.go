package main

import (
	"VOL/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义路由和处理函数
	router.POST("/k8s/command", handler.HandleK8sCommand)
	router.GET("/k8s/node", handler.Handler_get_node)
	router.GET("/k8s/vcjob/status", handler.HandleVcjobStatus)

	// 启动服务器
	router.Run(":8081")
}
