package main

import (
	auth "cloud_tinamic/kitex_gen/base/auth/authservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consulapi "github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"

	"net"
)

func main() {
	// Consul 注册中心的地址
	r, err := consul.NewConsulRegister("0.0.0.0:8500", consul.WithCheck(&consulapi.AgentServiceCheck{
		Interval:                       "7s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "1m",
	}))
	if err != nil {
		klog.Fatal("Failed to create Consul registry: ", err)
	}

	// 手动注册服务并指定服务名称
	svr := auth.NewServer(InitAuthService(),
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 8818, // 端口号
		}),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "auth.server",
		}),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal("auth server run failed error: %v", err)
	}
}
