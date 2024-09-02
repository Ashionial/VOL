package docker

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/registry"
	"os"
	"strings"
	"sync/atomic"
)

// 从系统中读取
// var authConfig = GetAuth()
var authConfig = registry.AuthConfig{
	Username:      "yjhknows",
	Password:      "dockerH50!258w",
	ServerAddress: "139.9.4.123:5000",
}

var authStr string
var counter int32

func Init() {
	authJson, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr = base64.URLEncoding.EncodeToString(authJson)

	counter = 0
}

func GetCount() int32 {
	atomic.AddInt32(&counter, 1)
	return counter
}

// 从系统中获取账号密码
func GetAuth() registry.AuthConfig {
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

	auth, exists := dockerConfig.Auths["https://index.docker.io/v1/"]
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

	authConfig := registry.AuthConfig{
		Username: authParts[0],
		Password: authParts[1],
	}

	fmt.Print("username: " + authConfig.Username)
	return authConfig
}

// 获取json文件结构的辅助结构体
type LocalDockerConfig struct {
	Auths map[string]Authenticate `json:"auths"`
}
type Authenticate struct {
	Auth string `json:"auth"`
}
