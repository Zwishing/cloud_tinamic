package main

import (
	"cloud_tinamic/kitex_gen/data/storage/storeservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	WorkflowClient "go.temporal.io/sdk/client"
)

func NewGeoServiceClient() (storeservice.Client, error) {
	c, err := storeservice.NewClient("geo", client.WithHostPorts("0.0.0.0:8089"))
	if err != nil {
		klog.Errorf("连接到storeservice失败: %s", err)
		return nil, err
	}
	return c, nil
}

func NewWorkflowClient() (WorkflowClient.Client, error) {
	c, err := WorkflowClient.Dial(WorkflowClient.Options{
		HostPort: WorkflowClient.DefaultHostPort,
	})
	if err != nil {
		klog.Fatalf("Unable to create client:%s", err)
		return nil, err
	}
	return c, nil
}
