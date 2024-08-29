package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVCJobHandler(c *gin.Context) {
	jobName := c.DefaultQuery("jobName", "all")
	namespace := c.DefaultQuery("namespace", "default")

	if jobName == "all" {
		// 查询所有vcjob的状态
		output, err := k8s.ExecuteCommandGetvcjobs(namespace)
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
		output, err := k8s.ExecuteCommandGetvcjob(jobName, namespace)
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
