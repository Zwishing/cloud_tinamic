package workflow

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"go.temporal.io/sdk/workflow"
	"time"
)

func VectorWorkflow(ctx workflow.Context, cloudOptimizedKey, srcPath, cloudOptimizedPath string) ([]byte, error) {

	// Get basket order.
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var vectorActivities *VectorActivities
	var cloudOptimizedFileSize int64
	err := workflow.ExecuteActivity(ctx, vectorActivities.ToGeoParquet, srcPath, cloudOptimizedPath).Get(ctx, &cloudOptimizedFileSize)
	if err != nil {
		return nil, err
	}

	var thumbnail []byte
	err = workflow.ExecuteActivity(ctx, vectorActivities.GenThumbnail, cloudOptimizedPath).Get(ctx, &thumbnail)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, vectorActivities.ToPostGIS, cloudOptimizedKey, cloudOptimizedPath).Get(ctx, nil)
	if err != nil {
		return nil, err
	}
	klog.Infof("")

	return thumbnail, nil

}
