package docker

import (
	"context"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"io"
)

func PushDockerImage(imageName string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	pushResponse, err := cli.ImagePush(ctx, imageName, image.PushOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return "", err
	}
	defer pushResponse.Close()

	_, err = io.Copy(io.Discard, pushResponse)
	if err != nil {
		return "Push Failed", err
	}

	return "Push Success", nil
}
