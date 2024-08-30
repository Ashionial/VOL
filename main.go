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
	router.POST("/docker", handler.DockerHandler)
	router.POST("/exec", handler.CmdHandler)
	router.POST("/k8s/command", handler.K8sCommandHandler)
	router.GET("/k8s/node", handler.GetNodeHandler)
	router.GET("/k8s/vcjob", handler.GetVCJobHandler)
	router.GET("/k8s/pod", handler.GetPodHandler)
	router.GET("/k8s/log", handler.LogHandler)
	// 启动服务器
	router.Run(":8081")
}
