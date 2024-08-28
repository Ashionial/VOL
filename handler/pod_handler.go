package handler

import (
	"VOL/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePodStatus(c *gin.Context) {
	var cmd K8sCommand

	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if cmd.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pod_name is required"})
		return
	}

	status, err := k8s.GetPodStatus(cmd.Name, cmd.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}
