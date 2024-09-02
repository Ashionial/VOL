package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os/exec"
)

func BuildDockerImage(imageName string, dockerfile string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	tarBag, err := CreateTarArchive([]byte(dockerfile))
	if err != nil {
		return "", err
	}

	buildOptions := types.ImageBuildOptions{
		Context:    tarBag,
		Dockerfile: "Dockerfile",
		Tags:       []string{imageName},
		Remove:     true,
	}

	response, err := cli.ImageBuild(ctx, tarBag, buildOptions)
	if err != nil {
		return "Failed to build", err
	}
	defer response.Body.Close()

	return imageName, nil
}

func CreateTarArchive(dockerfile []byte) (*bytes.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	header := &tar.Header{
		Name: "Dockerfile",
		Size: int64(len(dockerfile)),
	}

	if err := tw.WriteHeader(header); err != nil {
		return nil, err
	}
	if _, err := tw.Write(dockerfile); err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}

func BuildImageByFile(imageName string, contextPath string) (string, error) {
	buildCmd := exec.Command("docker", "build", "-t", imageName, contextPath)
	buildOutput, err := buildCmd.CombinedOutput()
	if err != nil {
		return string(buildOutput), err
		//return "", err
	}

	// 执行docker push命令
	result, err := PushDockerImage(imageName)
	if err != nil {
		return string(buildOutput) + "\n" + result, err
		//return "", err
	}

	// 返回成功信息
	return string(buildOutput) + "\n" + result, nil
}
