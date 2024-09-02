package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
)

func ManualHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if form == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "form is nil",
		})
		return
	}

	files_names := make([]string, 0)
	for _, file := range form.File["file"] {
		if file == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "file is nil",
			})
			return
		}
		c.SaveUploadedFile(file, "./file/"+file.Filename)
		files_names = append(files_names, file.Filename)
	}

	for _, cmd := range form.Value["cmd"] {
		req := strings.Split(cmd, " ")
		if len(req) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid cmd",
			})
			return
		}
	}
	outputs := ""
	for _, cmd := range form.Value["cmd"] {
		req := strings.Split(cmd, " ")
		for idx, request := range req {
			for _, name := range files_names {
				if request == name {
					req[idx] = "./file/" + req[idx]
					break
				}
			}
		}
		command := exec.Command(req[0], req[1:]...)
		output, err := command.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		outputs += "\n" + string(output)
	}
	encoded_output := base64.StdEncoding.EncodeToString([]byte(outputs))

	c.JSON(http.StatusOK, gin.H{
		"error":  "",
		"output": encoded_output,
	})
}
