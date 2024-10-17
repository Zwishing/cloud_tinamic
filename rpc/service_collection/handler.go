package main

import (
	source2 "cloud_tinamic/kitex_gen/data/source"
	source "cloud_tinamic/kitex_gen/data/source/sourceservice"
	"cloud_tinamic/kitex_gen/data/storage"
	"cloud_tinamic/kitex_gen/data/storage/storeservice"
	"cloud_tinamic/kitex_gen/map/processor"
	"cloud_tinamic/kitex_gen/map/processor/mapprocessorservice"
	collection "cloud_tinamic/kitex_gen/service/collection"
	"cloud_tinamic/pkg/errors"
	"cloud_tinamic/pkg/util"
	"cloud_tinamic/rpc/service_collection/repo"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// ServiceCollectionImpl implements the last ServiceCollection defined in the IDL.
type ServiceCollectionImpl struct {
	mapProcessorClient    mapprocessorservice.Client
	geoClient             storeservice.Client
	serviceCollectionRepo repo.ServiceCollectionRepo
	sourceClient          source.Client
}

func NewServiceCollectionImpl(mapProcessor mapprocessorservice.Client,
	geoStore storeservice.Client, serviceRepo repo.ServiceCollectionRepo,
	sourceService source.Client) *ServiceCollectionImpl {
	return &ServiceCollectionImpl{
		mapProcessorClient:    mapProcessor,
		geoClient:             geoStore,
		serviceCollectionRepo: serviceRepo,
		sourceClient:          sourceService,
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
func (s *ServiceCollectionImpl) Publish(ctx context.Context, req *collection.PublishRequest) (err error) {
	// 查询出数据的路径通过key
	path, err := s.sourceClient.GetUnifiedSourcePath(ctx, req.SourceKey)
	if err != nil {
		err = errors.Kerrorf(errors.DatabaseErrorCode, "获取源路径失败: %v", err)
		klog.Error(err)
		return err
	}

	path = "http://39.101.164.253:9000" + path

	var thumbnailCh = make(chan []byte, 1)

	// 启动 goroutine 处理缩略图生成
	go func() {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel() // 确保在函数结束时调用取消函数

		// 发起请求生成缩略图
		resp, err := s.mapProcessorClient.VectorThumbnail(timeoutCtx, &processor.VectorThumbnailRequest{
			FilePath: path, // 使用从源获取的实际路径
			Width:    300,
			Height:   300,
		})
		if err != nil {
			err = errors.Kerrorf(errors.InternalServerCode, "请求缩略图失败: %v", err)
			klog.Error(err)
			thumbnailCh <- nil
			return
		}
		thumbnailCh <- resp.GetThumbnail()
	}()

	title := util.GetFileName(path, false)
	switch req.SourceCategory {
	case source2.SourceCategory_VECTOR:
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel() // 确保在函数结束时调用取消函数
		_, err = s.geoClient.VectorStorage(timeoutCtx, &storage.StoreRequest{
			Schema: "vector",
			Table:  req.SourceKey,
			Name:   title,
			Url:    path,
			Ext:    util.GetFileExtension(path, false),
		})
		if err != nil {
			err = errors.Kerrorf(errors.DatabaseErrorCode, "向量存储失败: %v", err)
			klog.Error(err)
			return err
		}
	default:
		err = errors.Kerrorf(errors.InvalidInputCode, "不支持的源类别: %v", req.SourceCategory)
		klog.Error(err)
		return err
	}

	serviceKey := util.UuidV4()
	// 添加记录
	_, err = s.serviceCollectionRepo.AddCollection(req.SourceKey, serviceKey, title)
	if err != nil {
		err = errors.Kerrorf(errors.DatabaseErrorCode, "添加记录失败: %v", err)
		klog.Error(err)
		return err
	}

	go func() {
		thumbnail := <-thumbnailCh
		if thumbnail != nil {
			s.UpdateThumbnail(ctx, serviceKey, thumbnail)
		}
	}()

	return nil
}

func (s *ServiceCollectionImpl) UpdateThumbnail(ctx context.Context, serviceKey string, thumbnail []byte) error {
	_, err := s.serviceCollectionRepo.UpdateThumbnail(serviceKey, thumbnail)
	if err != nil {
		klog.Errorf("更新缩略图失败: %v", err)
		return err
	}
	klog.Infof("缩略图更新成功: %s", serviceKey)
	return nil
}
