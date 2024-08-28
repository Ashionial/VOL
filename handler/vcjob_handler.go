package handler

import (
	"VOL/k8s"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleVcjobStatus(c *gin.Context) {
	var cmd K8sCommand
	if cmd.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vcjob_name is required"})
		return
	}
	status, err := k8s.GetVcjobStatus(cmd.Name, cmd.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}
