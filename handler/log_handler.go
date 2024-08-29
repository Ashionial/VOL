package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogHandler(c *gin.Context) {
	logname := c.Query("logName")
	output, err := k8s.ExecuteCommandLog(logname)
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