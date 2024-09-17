package main

import (
	user "cloud_tinamic/kitex_gen/base/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
)

func main() {
	svr := user.NewServer(InitUserService(),
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 8810, // 端口号
		}))

	err := svr.Run()

	if err != nil {
		klog.Fatal("user server run failed error: %v", err)
	}
}
