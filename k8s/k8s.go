package k8s

import (
	"fmt"
	"os/exec"
)

// ExecuteCommand 执行 k8s 命令
func ExecuteCommand(action, resource, name, namespace string) (string, error) {
	// 示例：使用 kubectl 执行命令
	cmd := exec.Command("kubectl", action, resource, name, "-n", namespace)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询所有节点状态
func ExecuteCommand_getnodes() (string, error) {
	// 示例：使用 kubectl 执行命令
	cmd := exec.Command("kubectl", "get", "node")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询具体节点状态
func ExecuteCommand_getnode(node_name string) (string, error) {
	// 示例：使用 kubectl 执行命令
	cmd := exec.Command("kubectl", "discribe", "node", node_name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询具体vcjob状态
func GetVcjobStatus(name, namespace string) (string, error) {
	cmd := exec.Command("kubectl", "get", "vcjob", name, "-n", namespace, "-o", "jsonpath={.status.conditions[0].type}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get vcjob status: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}
