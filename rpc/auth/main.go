package main

import (
	auth "cloud_tinamic/kitex_gen/base/auth/authservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"

	consulapi "github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"

	"net"
)

func main() {
	r, err := createConsulRegistry()
	if err != nil {
		klog.Fatal("Failed to create Consul registry: ", err)
	}

	svr := createAuthServer("auth.server", 8818, r)

	err = svr.Run()

	if err != nil {
		klog.Fatal("auth server run failed error: ", err)
	}
}

func createConsulRegistry() (registry.Registry, error) {
	//39.101.164.253:8500
	return consul.NewConsulRegister("0.0.0.0:8500", consul.WithCheck(&consulapi.AgentServiceCheck{
		TCP:                            "0.0.0.0:8818",
		Interval:                       "7s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "1m",
	}))
}

func createAuthServer(serviceName string, port int, r registry.Registry) server.Server {
	return auth.NewServer(InitAuthService(),
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: port,
		}),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName,
		}),
	)
}
