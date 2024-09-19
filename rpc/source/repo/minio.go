package repo

import (
	conf "cloud_tinamic/config"
	"cloud_tinamic/pkg/storage"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

var (
	m         *storage.Storage
	minioOnce sync.Once
)

func NewMinio() *storage.Storage {
	minioOnce.Do(func() {
		cfg := conf.GetConfigInstance()
		minioConfig := storage.NewMinioConfig(
			storage.WithBucket(cfg.GetString("storage.minio.bucket")),
			storage.WithEndpoint(cfg.GetString("storage.minio.endpoint")),
			storage.WithRegion(cfg.GetString("storage.minio.region")),
			storage.WithToken(cfg.GetString("storage.minio.token")),
			storage.WithSecure(cfg.GetBool("storage.minio.secure")),
			storage.WithCredentials(
				cfg.GetString("storage.minio.accessKey"),
				cfg.GetString("storage.minio.secretKey")),
		)
		var err error
		m, err = storage.NewStorage(minioConfig)
		if err != nil {
			klog.Fatal("Failed Connect to Minio ", err)
		}
		klog.Infof("Connected to minio @ '%s' bucket '%s'", minioConfig.Endpoint, minioConfig.Bucket)
	})
	return m
}
