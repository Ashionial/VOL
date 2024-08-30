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
func ExecuteCommandGetNodes() (string, error) {
	// 示例：使用 kubectl 执行命令
	cmd := exec.Command("kubectl", "get", "node")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询具体节点状态
func ExecuteCommandGetNode(node_name string) (string, error) {
	// 示例：使用 kubectl 执行命令
	cmd := exec.Command("kubectl", "describe", "node", node_name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询所有vcjob的状态
func ExecuteCommandGetvcjobs(namespace string) (string, error) {
	cmd := exec.Command("kubectl", "get", "vcjob", "-n", namespace, "-o", "wide")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get all VCJobs: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询具体vcjob的状态
func ExecuteCommandGetvcjob(jobName, namespace string) (string, error) {
	cmd := exec.Command("kubectl", "get", "vcjob", jobName, "-n", namespace, "-o", "jsonpath={.status.conditions[0].type}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get VCJob status: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}

// 查询所有pod状态
func ExecuteCommandGetpods(namespace string) (string, error) {
	// 使用 kubectl 获取所有 Pod 的状态
	cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-o", "wide")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get all pods: %v, output: %s", err, string(output))
	}
	return string(output), nil
}

// 查询具体pod状态
func ExecuteCommandGetpod(podName, namespace string) (string, error) {
	// 使用 kubectl 获取指定 Pod 的状态
	cmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.phase}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get pod status: %v, output: %s", err, string(output))
	}
	return string(output), nil
}

func ExecuteCommandLog(podName string, namespace string) (string, error) {
	cmd := exec.Command("kubectl", "logs", podName, "--previous", "-n", namespace)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, string(output)+" "+cmd.String()+"\n")
	}
	return string(output), nil
}
