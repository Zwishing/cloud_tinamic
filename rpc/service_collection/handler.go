package main

import (
	"cloud_tinamic/kitex_gen/base"
	Source "cloud_tinamic/kitex_gen/data/source"
	SourceService "cloud_tinamic/kitex_gen/data/source/sourceservice"
	"cloud_tinamic/kitex_gen/data/storage"
	"cloud_tinamic/kitex_gen/data/storage/storeservice"
	"cloud_tinamic/kitex_gen/map/processor"
	"cloud_tinamic/kitex_gen/map/processor/mapprocessorservice"
	collection "cloud_tinamic/kitex_gen/service/collection"
	"cloud_tinamic/pkg"
	"cloud_tinamic/pkg/errors"
	"cloud_tinamic/pkg/util"
	"cloud_tinamic/rpc/service_collection/repo"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// ServiceCollectionImpl implements the last ServiceCollection defined in the IDL.
type ServiceCollectionImpl struct {
	MapProcessorClient    mapprocessorservice.Client
	GeoClient             storeservice.Client
	ServiceCollectionRepo repo.ServiceCollectionRepo
	SourceClient          SourceService.Client
}

func NewServiceCollectionImpl(mapProcessor mapprocessorservice.Client,
	geoStore storeservice.Client, serviceRepo repo.ServiceCollectionRepo,
	sourceService SourceService.Client) *ServiceCollectionImpl {
	return &ServiceCollectionImpl{
		MapProcessorClient:    mapProcessor,
		GeoClient:             geoStore,
		ServiceCollectionRepo: serviceRepo,
		SourceClient:          sourceService,
	}
}

// GetCollections implements the ServiceCollection interface.
func (s *ServiceCollectionImpl) GetCollections(ctx context.Context, pageSize int64, page int64) (resp *collection.GetCollectionsResponse, err error) {
	// TODO: Your code here...
	return
}

// AddCollection implements the ServiceCollection interface.
func (s *ServiceCollectionImpl) AddCollection(ctx context.Context, sourceKey string, title string) (resp *collection.AddCollectionResponse, err error) {
	// TODO: Your code here...
	return
}

// Publish implements the ServiceCollection interface.
func (s *ServiceCollectionImpl) Publish(ctx context.Context, req *collection.PublishRequest) (resp *collection.PublishResponse, err error) {
	resp = &collection.PublishResponse{
		Base: base.NewBaseResp(),
	}
	// 查询出数据的路径通过key
	cloudOptimizedKeyPath, err := s.SourceClient.GetCloudOptimizedSourcePath(ctx, req.SourceKey)
	if err != nil {
		err = errors.Kerrorf(errors.DatabaseErrorCode, "获取源路径失败: %v", err)
		klog.Error(err)
		resp.Base.Code = base.Code_NOT_FOUND
		resp.Base.Msg = err.Error()
		return resp, err
	}

	var serviceKeys []string

	for _, keyPathMap := range cloudOptimizedKeyPath {

		var thumbnailCh = make(chan []byte, 1)

		// 启动 goroutine 处理缩略图生成
		go func() {
			timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel() // 确保在函数结束时调用取消函数

			// 发起请求生成缩略图
			resp, err := s.MapProcessorClient.VectorThumbnail(timeoutCtx, &processor.VectorThumbnailRequest{
				CloudOptimizedPath:       keyPathMap["path"], // 使用从源获取的实际路径
				CloudOptimizedBucketName: pkg.CloudOptimizedSourceBucketName,
				Width:                    300,
				Height:                   300,
			})
			if err != nil {
				err = errors.Kerrorf(errors.InternalServerCode, "请求缩略图失败: %v", err)
				klog.Error(err)
				thumbnailCh <- nil
				return
			}
			thumbnailCh <- resp.GetThumbnail()
		}()

		switch req.SourceCategory {
		case Source.SourceCategory_VECTOR:
			timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel() // 确保在函数结束时调用取消函数
			_, err = s.GeoClient.VectorToPGStorage(timeoutCtx, &storage.VectorToPGStorageRequest{
				Schema:                   pkg.VectorSchema,  // 存储矢量数据的schema
				Table:                    keyPathMap["key"], // cloud_optimized_key
				Name:                     req.ServiceName,
				CloudOptimizedBucketName: pkg.CloudOptimizedSourceBucketName,
				CloudOptimizedPath:       keyPathMap["path"],
			})
			if err != nil {
				err = errors.Kerrorf(errors.DatabaseErrorCode, "vector存储失败: %v", err)
				klog.Error(err)
				resp.Base.Code = base.Code_FAIL
				resp.Base.Msg = err.Error()
				return resp, err
			}
		default:
			err = errors.Kerrorf(errors.InvalidInputCode, "不支持的源类别: %v", req.SourceCategory)
			klog.Error(err)
			resp.Base.Code = base.Code_FAIL
			resp.Base.Msg = err.Error()
			return resp, err
		}

		serviceKey := util.UuidV4()

		serviceKeys = append(serviceKeys, serviceKey)
		// 添加记录
		_, err = s.ServiceCollectionRepo.AddCollection(keyPathMap["key"], serviceKey, req.ServiceName)
		if err != nil {
			err = errors.Kerrorf(errors.DatabaseErrorCode, "添加记录失败: %v", err)
			klog.Error(err)
			resp.Base.Code = base.Code_SERVER_ERROR
			resp.Base.Msg = err.Error()
			return resp, err
		}

		thumbnail := <-thumbnailCh
		if thumbnail != nil {
			err := s.UpdateThumbnail(ctx, serviceKey, thumbnail)
			if err != nil {
				klog.Errorf("未能更新 %s 缩率图: %v", serviceKey, err)
			}
		}
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.ServiceKeys = serviceKeys
	return resp, nil
}

func (s *ServiceCollectionImpl) UpdateThumbnail(ctx context.Context, serviceKey string, thumbnail []byte) error {
	_, err := s.ServiceCollectionRepo.UpdateThumbnail(serviceKey, thumbnail)
	if err != nil {
		klog.Errorf("更新缩略图失败: %v", err)
		return err
	}
	klog.Infof("缩略图更新成功: %s", serviceKey)
	return nil
}
