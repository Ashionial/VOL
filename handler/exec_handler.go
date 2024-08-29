package handler

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
)

func CmdHandler(c *gin.Context) {
	cmdStr := c.PostForm("cmd")

	if !strings.HasPrefix(cmdStr, "kubectl") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "kubectl command not found",
			"cmd":   cmdStr,
		})
		return
	}

	req := strings.Split(cmdStr, " ")
	cmd := exec.Command(req[0], req[1:]...)
	output, err := cmd.CombinedOutput()
	encodedOutput := base64.StdEncoding.EncodeToString(output)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"output": encodedOutput,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  "",
		"output": encodedOutput,
	})
}
