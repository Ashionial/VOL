package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetVCJobStatus(c *gin.Context) {
	jobName := c.DefaultQuery("jobName", "all")
	namespace := c.DefaultQuery("namespace", "default")

	if jobName == "all" {
		// 查询所有vcjob的状态
		output, err := k8s.ExecuteCommand_getAllVCJobs(namespace)
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
		// 查询指定vcjob的状态
		output, err := k8s.ExecuteCommand_getVCJob(jobName, namespace)
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
