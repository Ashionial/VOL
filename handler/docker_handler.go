package handler

import (
	"encoding/base64"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

// 因为只有dockerfile所以只能使用官方库中的内容，还要进一步改进以支持文件的上传
// 如果一定要使用json格式，可以考虑文件转化为Base64编码嵌入json中，但是json文件会变得很大
// 也可以直接传文件
func DockerHandler(c *gin.Context) {
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

	imageName := make([]string, 0)
	for _, file := range form.File["file"] {
		if file == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "file is nil",
			})
			return
		}
		c.SaveUploadedFile(file, "./file/"+"139.9.4.123:5000/yjhknows/"+file.Filename)
		imageName = append(imageName, file.Filename)
	}

	outputs := ""
	for _, cmd := range form.Value["cmd"] {
		req := strings.Split(cmd, " ")
		for idx, request := range req {
			for _, name := range imageName {
				if request == name {
					req[idx] = "./file/" + req[idx]
					break
				}
			}
		}

		if len(req) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid cmd",
			})
			return
		}

		command := exec.Command(req[0], req[1:]...)
		output, err := command.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		outputs += "\n" + string(output)
	}

	encodedOutput := base64.StdEncoding.EncodeToString([]byte(outputs))

	c.JSON(http.StatusOK, gin.H{
		"error":  "",
		"output": encodedOutput,
	})
}
