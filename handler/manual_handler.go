package handler

import (
	"VOL/docker"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ManualHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to parse multipart form",
		})
		return
	}
	if form == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "form is nil",
			"message": "Failed to parse multipart form",
		})
		return
	}

	path := "./file" + strconv.Itoa(int(docker.GetCount())) + "/"

	files_names := make([]string, 0)
	for _, file := range form.File["file"] {
		if file == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "file is nil",
			})
			return
		}
		c.SaveUploadedFile(file, path+file.Filename)
		files_names = append(files_names, file.Filename)
	}

	dockerfileContent := fmt.Sprintf(`
	FROM python:3.9-slim
	
	WORKDIR /app

	COPY . /app

	# RUN pip install --no-cache-dir -r requirements.txt
	`)

	command := c.PostForm("cmd")
	commands := strings.Split(command, " ")
	dockerfileContent += fmt.Sprintf(`CMD ["%s"`, commands[0])
	for _, cmd := range commands[1:] {
		dockerfileContent += fmt.Sprintf(`, "%s"`, cmd)
	}
	dockerfileContent += "]\n"

	err = os.WriteFile(path+"Dockerfile", []byte(dockerfileContent), os.ModePerm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to create docker file",
		})
		return
	}

	imageName := c.PostForm("imageName")
	imageName = "139.9.4.123:5000/yjhknows/" + strings.ToLower(imageName)
	buildOutput, err := docker.BuildImageByFile(imageName, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":       "",
		"buildOutput": buildOutput,
	})
}
