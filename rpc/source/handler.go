package main

import (
	"bytes"
	"cloud_tinamic/kitex_gen/base"
	source "cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/util"
	"cloud_tinamic/rpc/source/pack"
	"cloud_tinamic/rpc/source/repo"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// SourceServiceImpl implements the last service interface defined in the IDL.
type SourceServiceImpl struct {
	SourceRepo repo.SourceRepo
}

func NewSourceServiceImpl(repo repo.SourceRepo) *SourceServiceImpl {
	return &SourceServiceImpl{SourceRepo: repo}
}

// Upload implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) Upload(ctx context.Context, req *source.UploadRequest) (resp *source.UploadResponse, err error) {
	resp = &source.UploadResponse{
		Base: base.NewBaseResp(),
	}

	bucketName := strings.ToLower(req.SourceCategory.String())

	if req.Key == "" {
		req.Key, err = s.SourceRepo.GetHomeKeyBySourceCategory(req.SourceCategory)
	}

	path, err := s.SourceRepo.GetPathByKey(req.Key)
	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "获取路径失败"
		return resp, err
	}
	newPath := fmt.Sprintf("%s/%s", path, req.Name)
	Key := util.UuidV4()

	errChan := make(chan error, 1) // 通道大小可以设置为1
	var wg sync.WaitGroup
	wg.Add(2)
	doneChan := make(chan struct{}) // 新增一个通道用于表示任务完成

	// Goroutine 1: 上传文件到 Minio
	go func() {
		defer wg.Done()
		err := s.SourceRepo.UploadToMinio(bucketName, newPath, req.Name, bytes.NewReader(req.FileData), req.Size)
		if err != nil {
			errChan <- err
		}
	}()

	// Goroutine 2: 添加文件元数据到数据库
	go func() {
		defer wg.Done()
		success, addItemErr := s.SourceRepo.AddItem(req.SourceCategory, req.Key, &source.Item{
			Name:         req.Name,
			ItemType:     source.ItemType_FILE,
			Key:          Key,
			Size:         req.Size,
			ModifiedTime: time.Now().Unix(),
			Path:         newPath,
		})
		if addItemErr != nil || !success {
			errChan <- fmt.Errorf("添加到数据库失败: %v", addItemErr)
		}
	}()

	// 启动一个 goroutine 来等待所有任务完成
	go func() {
		wg.Wait()       // 等待所有 goroutine 完成
		close(doneChan) // 完成后关闭 doneChan
	}()

	// 处理 select 语句
	select {
	case err := <-errChan: // 任何错误会立即触发
		if strings.Contains(err.Error(), "添加到数据库失败") {
			//_ = s.SourceRepo.DeleteFromMinio(bucketName, newPath) // 保持一致性，删除已上传文件
			resp.Base.Code = base.Code_FAIL
			resp.Base.Msg = "添加到数据库失败"
		} else {
			resp.Base.Code = base.Code_FAIL
			resp.Base.Msg = "上传到 Minio 失败"
		}
		return resp, err
	case <-doneChan: // 如果所有 goroutine 都完成，没有错误，立即返回成功
		resp.Base.Code = base.Code_SUCCESS
		resp.Base.Msg = "数据上传成功"
		resp.Key = Key
		return resp, nil
	case <-time.After(300 * time.Second): // 超时处理
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "上传超时"
		return resp, fmt.Errorf("上传超时")
	}

}

// PresignedUpload implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) PresignedUpload(ctx context.Context, req *source.PresignedUploadResquest) (resp *source.PresignedUploadResponse, err error) {
	resp = source.NewPresignedUploadResponse()
	resp.SetBase(base.NewBaseResp())

	url, err := s.SourceRepo.PresignedUploadUrl(req.SourceCategory, req.Path, req.Name)
	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "无法生成上传URL"
		resp.Url = ""
		return
	}
	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "成功生成上传URL"
	resp.Url = url
	return
}

// GetNextItems NextItems implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) GetNextItems(ctx context.Context, req *source.GetItemsRequest) (resp *source.GetItemsResponse, err error) {
	resp = &source.GetItemsResponse{
		Base: base.NewBaseResp(),
	}

	items, err := s.SourceRepo.GetChildrenItemsByKey(req.Key)

	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "获取记录失败"
		return resp, err
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "返回成功"
	resp.Items = pack.Storages(items)
	resp.Key = req.Key
	return resp, nil
}

func (s *SourceServiceImpl) GetPreviousItems(ctx context.Context, req *source.GetItemsRequest) (resp *source.GetItemsResponse, err error) {
	resp = &source.GetItemsResponse{
		Base: base.NewBaseResp(),
	}
	items, err := s.SourceRepo.GetSiblingItemsByKey(req.Key)

	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "获取记录失败"
		return resp, err
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "返回成功"
	resp.Items = pack.Storages(items)
	if len(items) > 0 {
		resp.Key = items[0].ParentKey
	}
	return resp, nil

}

func (s *SourceServiceImpl) GetHomeItems(ctx context.Context, req *source.GetHomeItemsRequest) (resp *source.GetItemsResponse, err error) {
	resp = &source.GetItemsResponse{
		Base: base.NewBaseResp(),
	}
	key, items, err := s.SourceRepo.GetHomeItemsBySourceCategory(req.SourceCategory)

	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "获取记录失败"
		return resp, err
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "返回成功"
	resp.Items = pack.Storages(items)
	resp.Key = key
	return resp, nil
}

// CreateFolder implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) CreateFolder(ctx context.Context, req *source.CreateFolderRequest) (resp *source.AddItemResponse, err error) {
	resp = &source.AddItemResponse{
		Base: base.NewBaseResp(),
	}

	count, err := s.SourceRepo.GetCountByName(req.Key, req.Name, source.ItemType_FOLDER)
	if err != nil {
		resp.Base.Code = base.Code_SERVER_ERROR
		resp.Base.Msg = "获取文件夹数量失败"
		return
	}
	if count > 0 {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "同名文件夹已存在"
		return
	}

	folder := &source.Item{
		ParentKey:    req.Key,
		Name:         req.Name,
		ItemType:     source.ItemType_FOLDER,
		Key:          util.UuidV4(),
		Size:         0,
		ModifiedTime: time.Now().Local().Unix(),
		Path:         req.Path,
	}

	success, err := s.SourceRepo.AddItem(req.SourceCategory, req.Key, folder)
	if err != nil {
		resp.Base.Code = base.Code_SERVER_ERROR
		resp.Base.Msg = "创建文件夹失败"
		return resp, err
	}
	if !success {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "创建文件夹失败"
		return resp, err
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "创建文件夹成功"
	resp.Item = folder
	return resp, nil
}

// DeleteItems implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) DeleteItems(ctx context.Context, req *source.DeleteItemsRequest) (resp *source.DeleteItemsResponse, err error) {
	resp = &source.DeleteItemsResponse{
		Base: base.NewBaseResp(),
	}

	isDeleted, err := s.SourceRepo.DeleteItems(req.Keys)
	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "删除发生错误，删除失败"
		return resp, err
	}
	if !isDeleted {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "删除失败"
		return resp, nil
	}
	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "删除成功"
	return resp, nil
}

func (s *SourceServiceImpl) AddItem(ctx context.Context, req *source.AddItemRequest) (resp *source.AddItemResponse, err error) {
	resp = source.NewAddItemResponse()
	resp.SetBase(base.NewBaseResp())

	success, err := s.SourceRepo.AddItem(req.SourceCategory, req.CurrentFolder, req.Item)
	if err != nil {
		resp.Base.Code = base.Code_SERVER_ERROR
		resp.Base.Msg = "添加错误"
		return resp, err
	}
	if !success {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "添加失败"
		return resp, err
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "添加成功"
	return resp, nil
}

func (s *SourceServiceImpl) GetUnifiedSourcePath(ctx context.Context, key string) (string, error) {
	path, err := s.SourceRepo.GetPathByKey(key)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (s *SourceServiceImpl) GetSourcePath(ctx context.Context, key string) (string, error) {
	path, err := s.SourceRepo.GetPathByKey(key)
	if err != nil {
		return "", err
	}
	return path, nil
}
