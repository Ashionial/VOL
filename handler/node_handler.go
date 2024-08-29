package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNodeHandler(c *gin.Context) {
	username := c.DefaultQuery("username", "all")
	if username == "all" {
		// 执行 Kubernetes 命令
		output, err := k8s.ExecuteCommandGetNodes()
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
		output, err := k8s.ExecuteCommandGetNode(username)
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
