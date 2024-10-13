package main

import (
	vector "cloud_tinamic/kitex_gen/service/vector/vectorservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
)

func main() {
	svr := vector.NewServer(InitVectorService(),
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 8816, // 端口号
		}))

	err := svr.Run()

	if err != nil {
		klog.Fatal("vector server run failed error: %v", err)
	}
}
