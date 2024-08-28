package handler

import (
	"VOL/k8s"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler_get_node(c *gin.Context) {
	username := c.DefaultQuery("username", "all")
	if username == "all" {
		// 执行 Kubernetes 命令
		output, err := k8s.ExecuteCommand_getnodes()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "output": output})
			return
		}
		c.String(http.StatusOK, output)
	} else {
		output, err := k8s.ExecuteCommand_getnode(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "output": output})
			return
		}
		c.String(http.StatusOK, output)
	}

}
