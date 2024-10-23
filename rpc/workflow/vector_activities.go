package workflow

import (
	"cloud_tinamic/kitex_gen/base"
	"cloud_tinamic/kitex_gen/data/storage"
	"cloud_tinamic/kitex_gen/data/storage/storeservice"
	"cloud_tinamic/kitex_gen/map/processor"
	"cloud_tinamic/kitex_gen/map/processor/mapprocessorservice"
	"cloud_tinamic/pkg"
	"context"
)

type VectorActivities struct {
	MapProcessorClient mapprocessorservice.Client
	GeoClient          storeservice.Client
}

func NewVectorActivities(m mapprocessorservice.Client, geo storeservice.Client) *VectorActivities {
	return &VectorActivities{
		MapProcessorClient: m,
		GeoClient:          geo,
	}
}

func (v *VectorActivities) ToGeoParquet(srcPath string, destPath string) (int64, error) {
	resp, err := v.GeoClient.ToGeoParquetStorage(context.Background(), &storage.ToGeoParquetStorageRequest{
		SourceBucket: pkg.OriginalSourceBucketName,
		SourcePath:   srcPath,
		DestBucket:   pkg.CloudOptimizedSourceBucketName,
		DestPath:     destPath,
	})
	if err != nil {
		return 0, err
	}
	if resp.Base.Code != base.Code_SUCCESS {
		return 0, nil
	}
	return resp.Size, nil
}

func (v *VectorActivities) GenThumbnail(cloudOptimizedPath string) ([]byte, error) {
	resp, err := v.MapProcessorClient.VectorThumbnail(context.Background(), &processor.VectorThumbnailRequest{
		CloudOptimizedPath:       cloudOptimizedPath, // 使用从源获取的实际路径
		CloudOptimizedBucketName: pkg.CloudOptimizedSourceBucketName,
		Width:                    pkg.VectorThumbnailWidth,
		Height:                   pkg.VectorThumbnailHeight,
	})
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != base.Code_SUCCESS {
		return nil, nil
	}
	return resp.Thumbnail, nil
}

func (v *VectorActivities) ToPostGIS(CloudOptimizedKey, cloudOptimizedPath string) error {
	resp, err := v.GeoClient.VectorToPGStorage(context.Background(), &storage.VectorToPGStorageRequest{
		Schema:                   pkg.VectorSchema,  // 存储矢量数据的schema
		Table:                    CloudOptimizedKey, // cloud_optimized_key
		Name:                     "",
		CloudOptimizedBucketName: pkg.CloudOptimizedSourceBucketName,
		CloudOptimizedPath:       cloudOptimizedPath,
	})
	if err != nil {
		return err
	}
	if resp.Base.Code != base.Code_SUCCESS {
		return err
	}
	return nil
}
