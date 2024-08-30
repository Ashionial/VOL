package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogHandler(c *gin.Context) {
	//查询正在进行的pob的log
	podName := c.Query("podName")
	namespace := c.Query("namespace")
	output, err := k8s.ExecuteCommandLog(podName, namespace)
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
