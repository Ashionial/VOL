package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"os"
	"strings"
)

// 从系统中读取
var authConfig = GetAuth()

func PushDockerImage(imageName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	ctx := context.Background()
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	_, err = cli.ImagePush(ctx, imageName, image.PushOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return err
	}

	return nil
}

// 从系统中获取账号密码
func GetAuth() AuthConfig {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := homeDir + "/.docker/config.json"
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err) // 处理文件关闭错误
		}
	}()

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var dockerConfig LocalDockerConfig
	err = json.Unmarshal(data, &dockerConfig)
	if err != nil {
		panic(err)
	}

	auth, exists := dockerConfig.Auths["http://index.docker.io/v1/"]
	if !exists {
		panic("No auth config found")
	}

	decodedAuth, err := base64.StdEncoding.DecodeString(auth.Auth)
	if err != nil {
		panic(err)
	}

	authParts := strings.Split(string(decodedAuth), ":")
	if len(authParts) != 2 {
		panic("Invalid auth config format")
	}

	authConfig := AuthConfig{
		Username: authParts[0],
		Password: authParts[1],
	}

	return authConfig
}

// 获取json文件结构的辅助结构体
type LocalDockerConfig struct {
	Auths map[string]Authenticate `json:"auths"`
}
type Authenticate struct {
	Auth string `json:"auth"`
}

// types中找不到这个结构体，我直接复制了一份下来
// AuthConfig contains authorization information for connecting to a Registry
type AuthConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Auth     string `json:"auth,omitempty"`

	// Email is an optional value associated with the username.
	// This field is deprecated and will be removed in a later
	// version of docker.
	Email string `json:"email,omitempty"`

	ServerAddress string `json:"serveraddress,omitempty"`

	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	IdentityToken string `json:"identitytoken,omitempty"`

	// RegistryToken is a bearer token to be sent to a registry
	RegistryToken string `json:"registrytoken,omitempty"`
}
