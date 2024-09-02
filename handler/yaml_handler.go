package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetYamlHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is nil",
		})
		return
	}

	file_name := file.Filename

	c.SaveUploadedFile(file, "./file/"+file.Filename)

	output, err := k8s.ExecuteCommandYaml(file_name)

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
