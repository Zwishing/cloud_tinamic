package main

import (
	"cloud_tinamic/pkg"
	"cloud_tinamic/rpc/workflow"
	"cloud_tinamic/rpc/workflow/wire"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort, // ip:port
	})

	if err != nil {
		klog.Fatalf("Unable to create client: %v", err)
	}
	defer c.Close()
	w := worker.New(c, pkg.VectorProcessingTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.VectorWorkflow)

	activities, err := wire.InitVectorActivities()

	if err != nil {
		klog.Fatalf("Unable to start vector activities: %v", err)
	}

	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())

	if err != nil {
		klog.Fatalf("Unable to start worker: %v", err)
	}
}
