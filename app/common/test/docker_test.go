package main

import (
	"context"
	"github.com/docker/docker/api/types"
	container2 "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"testing"
)

func Test(t *testing.T) {
	// 连接到Docker服务器
	ctx := context.Background()
	dockerClient, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	// 创建一个容器
	container, _ := dockerClient.ContainerCreate(ctx, &container2.Config{
		Image: "test",
	}, nil, nil, nil, "test1")

	// 启动容器
	dockerClient.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})

}
