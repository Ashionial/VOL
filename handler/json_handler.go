package handler

import (
	"VOL/k8s"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type K8sCommand struct {
	Action    string `json:"action" binding:"required"`
	Resource  string `json:"resource" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Namespace string `json:"namespace"`
}

func HandleK8sCommand(c *gin.Context) {
	var cmd K8sCommand

	// 解析 JSON
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 执行 Kubernetes 命令
	output, err := k8s.ExecuteCommand(cmd.Action, cmd.Resource, cmd.Name, cmd.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": output})
		return
	}

	c.JSON(http.StatusOK, gin.H{"output": output})
}

func HandlerGetNode(c *gin.Context) {
	username := c.DefaultQuery("username", "all")
	if username == "all" {
		// 执行 Kubernetes 命令
		output, err := k8s.ExecuteCommand_getnodes()
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
		output, err := k8s.ExecuteCommand_getnode(username)
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

func HandlerGetVCJobStatus(c *gin.Context) {
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

func HandlerGetPodStatus(c *gin.Context) {
	podName := c.DefaultQuery("podName", "all")
	namespace := c.DefaultQuery("namespace", "default")

	if podName == "all" {
		// 查询所有pod状态的命令
		output, err := k8s.ExecuteCommand_getAllPods(namespace)
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
		// 查询指定pod的状态
		output, err := k8s.ExecuteCommand_getPod(podName, namespace)
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
