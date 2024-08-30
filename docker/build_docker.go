package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
		Dockerfile: "dockerfile",
		Tags:       []string{imageName},
		Remove:     true,
	}

	response, err := cli.ImageBuild(ctx, tarBag, buildOptions)
	if err != nil {
		return "", err
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
