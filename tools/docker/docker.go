package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

func main() {
	imagePtr := flag.String("i", "", "-i image")
	flag.Parse()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := listAllContainers(cli)
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if strings.Contains(container.Image, *imagePtr) {
			err = connectToContainer(cli, container.ID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Connected to container successfully.")
			}
		}
	}

	//// 容器 ID 或名称
	//containerID := "your-container-id"
	//
	//// 新的端口映射规则
	//newPortMapping := "9090:80"
	//
	//// 调用函数修改端口映射
	//err = updatePortMapping(cli, containerID, newPortMapping)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Port mapping updated successfully.")
	//}
}

func connectToContainer(cli *client.Client, containerID string) error {
	// 使用 exec 创建一个连接到容器的 exec 实例
	execCreateResp, err := cli.ContainerExecCreate(context.Background(), containerID, types.ExecConfig{
		Cmd:          []string{"/bin/sh"},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	})
	if err != nil {
		return err
	}

	// 启动 exec 实例
	execStartResp, err := cli.ContainerExecAttach(context.Background(), execCreateResp.ID, types.ExecStartCheck{Tty: true})
	if err != nil {
		return err
	}

	// 在执行完命令后关闭 exec 实例
	defer execStartResp.Close()

	return nil
}

func listAllContainers(cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func updatePortMapping(cli *client.Client, containerID string, newPortMapping string) error {
	// 获取容器的当前配置
	container, _, err := cli.ContainerInspectWithRaw(context.Background(), containerID, false)
	if err != nil {
		return err
	}

	// 修改端口映射规则
	container.HostConfig.PortBindings["80/tcp"] = []types.PortBinding{{HostIP: "0.0.0.0", HostPort: "9090"}}

	// 提交修改后的配置
	updateConfig := types.ContainerUpdateConfig{
		HostConfig: container.HostConfig,
	}
	_, err = cli.ContainerUpdate(context.Background(), containerID, updateConfig)
	if err != nil {
		return err
	}

	return nil
}
