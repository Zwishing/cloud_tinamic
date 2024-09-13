package main

import (
	auth "cloud_tinamic/kitex_gen/base/auth/authservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
)

func main() {
	svr := auth.NewServer(InitAuthService(),
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 8811, // 端口号
		}))

	err := svr.Run()

	if err != nil {
		klog.Error("user server run failed error: %v", err)
	}
}
