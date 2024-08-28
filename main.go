package main

import (
	"VOL/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// ����·�ɺʹ�����
	router.POST("/k8s/command", handler.HandleK8sCommand)
	router.GET("/k8s/node", handler.Handler_get_node)
	router.GET("/k8s/vcjob/status", handler.HandleVcjobStatus)
	router.GET("/k8s/pod_status", handler.HandlePodStatus) // ����·��

	// ����������
	router.Run(":8081")
}
