package handler

import (
	"VOL/docker"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 因为只有dockerfile所以只能使用官方库中的内容，还要进一步改进以支持文件的上传
// 如果一定要使用json格式，可以考虑文件转化为Base64编码嵌入json中，但是json文件会变得很大
// 也可以直接传文件
func DockerHandler(c *gin.Context) {
	imageName := c.PostForm("imageName")
	dockerfileContent := c.PostForm("dockerfile")

	imageName, err := docker.BuildDockerImage(imageName, dockerfileContent)
	imageName = "139.9.4.123:5000/yjhknows/" + strings.ToLower(imageName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := docker.PushDockerImage(imageName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image":   imageName,
		"message": "Docker image built and pushed successfully",
	})
}
