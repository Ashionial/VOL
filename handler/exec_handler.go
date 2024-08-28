package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
)

func HandleCmd(c *gin.Context) {
	cmdStr := c.Query("cmd")

	if !strings.HasPrefix(cmdStr, "kubectl") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "kubectl command not found",
		})
		return
	}

	req := strings.Split(cmdStr, " ")
	cmd := exec.Command(req[0], req[1:]...)
	output, err := cmd.CombinedOutput()
	encoded_output := base64.StdEncoding.EncodeToString(output)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"output": encoded_output,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  "",
		"output": encoded_output,
	})
}
