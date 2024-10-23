package workflow

import (
	"cloud_tinamic/kitex_gen/data/storage/storeservice"
	"cloud_tinamic/kitex_gen/map/processor/mapprocessorservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func NewMapProcessorClient() (mapprocessorservice.Client, error) {
	c, err := mapprocessorservice.NewClient("map", client.WithHostPorts("0.0.0.0:9090"))
	if err != nil {
		klog.Errorf("连接到mapprocessorservice失败: %s", err)
		return nil, err
	}
	return c, nil
}

func NewGeoServiceClient() (storeservice.Client, error) {
	c, err := storeservice.NewClient("geo", client.WithHostPorts("0.0.0.0:8089"))
	if err != nil {
		klog.Errorf("连接到storeservice失败: %s", err)
		return nil, err
	}
	return c, nil
}
