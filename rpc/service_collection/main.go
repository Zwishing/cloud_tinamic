package main

import (
	collection "cloud_tinamic/kitex_gen/service/collection/servicecollection"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	server, err := InitServiceCollection()
	if err != nil {
		klog.Errorf("初始化服务集合失败: %v", err)
		return
	}
	svr := collection.NewServer(server)
	err = svr.Run()

	if err != nil {
		klog.Errorf("运行服务器失败: %v", err)
		return
	}
}
