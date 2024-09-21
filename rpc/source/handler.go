package main

import (
	"bytes"
	"cloud_tinamic/kitex_gen/base"
	source "cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/util"
	"cloud_tinamic/rpc/source/model"
	"cloud_tinamic/rpc/source/pack"
	"cloud_tinamic/rpc/source/repo"
	"context"
	"fmt"
	"strings"
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
	path, err := s.SourceRepo.GetPathByKey(req.Key)
	newPath := fmt.Sprintf("%s/%s", path, req.Name)
	Key := util.UuidV4()

	errChan := make(chan error, 2)
	var success bool

	go func() {
		// 上传到minio
		// TODO 后续可以使用hash对重复的数据就可以不用在上传
		err := s.SourceRepo.UploadToMinio(bucketName, newPath, req.Name, bytes.NewReader(req.FileData), req.Size)
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		// 添加到数据记录
		var addItemErr error
		success, addItemErr = s.SourceRepo.AddItem(req.SourceCategory, req.Key, &source.Item{
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

	select {
	case err := <-errChan:
		if strings.Contains(err.Error(), "添加到数据库失败") {
			// TODO 为了保持数据一致性，需要删除已经上传到minio的数据
			resp.Base.Code = base.Code_FAIL
			resp.Base.Msg = "添加到数据库失败"
		} else {
			resp.Base.Code = base.Code_FAIL
			resp.Base.Msg = "上传到minio失败"
		}
		return resp, err
	case <-time.After(300 * time.Second):
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "上传超时"
		return resp, fmt.Errorf("上传超时")
	default:
		resp.Base.Code = base.Code_SUCCESS
		resp.Base.Msg = "数据上传成功"
		resp.Key = Key
		return resp, nil
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

// GetItem implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) GetItem(ctx context.Context, req *source.GetItemRequest) (resp *source.GetItemResponse, err error) {
	resp = &source.GetItemResponse{
		Base: base.NewBaseResp(),
	}
	var items []*model.Storage
	var key string
	if req.Key == "" {
		key, items, err = s.SourceRepo.GetTopItemsBySourceCategory(req.SourceCategory)
	} else {
		key = req.Key
		items, err = s.SourceRepo.GetItemByKey(req.Key)
	}

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
		ModifiedTime: time.Now().Unix(),
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

// DeleteItem implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) DeleteItem(ctx context.Context, req *source.DeleteItemRequest) (resp *source.DeleteItemResponse, err error) {
	// TODO: Your code here...
	return
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
