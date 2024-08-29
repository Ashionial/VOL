package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPodHandler(c *gin.Context) {
	podName := c.DefaultQuery("podName", "all")
	namespace := c.DefaultQuery("namespace", "default")

	if podName == "all" {
		// 查询所有pod状态的命令
		output, err := k8s.ExecuteCommandGetpods(namespace)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"output": base64.StdEncoding.EncodeToString([]byte(output)),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"output": base64.StdEncoding.EncodeToString([]byte(output)),
			"error":  "",
		})
	} else {
		// 查询指定pod的状态
		output, err := k8s.ExecuteCommandGetpod(podName, namespace)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"output": base64.StdEncoding.EncodeToString([]byte(output)),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"output": base64.StdEncoding.EncodeToString([]byte(output)),
			"error":  "",
		})
	}
}
