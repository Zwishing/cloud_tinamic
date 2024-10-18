package main

import (
	source "cloud_tinamic/kitex_gen/data/source/sourceservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
)

func main() {

	service, err := InitSourceService()
	svr := source.NewServer(
		service,
		server.WithServiceAddr(
			&net.TCPAddr{
				IP:   net.ParseIP("0.0.0.0"),
				Port: 8813, // 端口号
			}))

	// 启动nsq
	client := InitNsqClient()
	// 关闭nsq
	defer client.Stop()

	err = svr.Run()
	if err != nil {
		klog.Fatal("user server run failed error: %v", err)
	}
}
