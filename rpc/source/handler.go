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
	resp = source.NewUploadResponse()
	resp.SetBase(base.NewBaseResp())

	bucketName := strings.ToLower(req.SourceCategory.String())
	path, err := s.SourceRepo.GetPathById(req.ParentId)
	newPath := fmt.Sprintf("%s/%s", path, req.Name)
	// 上传到minio
	// TODO 后续可以使用hash对重复的数据就可以不用在上传
	err = s.SourceRepo.UploadToMinio(bucketName, newPath, req.Name, bytes.NewReader(req.FileData), req.Size)
	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "上传到minio失败"
		return
	}
	// 上传成功后添加到数据记录
	sourceId := util.UuidV4()
	success, err := s.SourceRepo.AddItem(req.SourceCategory, req.ParentId, &source.Item{
		Name:         req.Name,
		ItemType:     source.ItemType_FILE,
		Key:          sourceId,
		Size:         req.Size,
		ModifiedTime: time.Now().Unix(),
		Path:         newPath,
	})
	if err != nil || success == false {
		// TODO 为了保持数据一致性，需要删除已经上传到minio的数据
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "添加到数据库失败"
		return
	}

	resp.SourceId = sourceId
	return
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
	resp = source.NewGetItemResponse()
	resp.SetBase(base.NewBaseResp())

	items, err := s.SourceRepo.GetItemById(req.Key)
	if err != nil {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "获取记录失败"
		return
	}
	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "返回成功"
	resp.Items = pack.Storages(items)
	return
}

// CreateFolder implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) CreateFolder(ctx context.Context, req *source.CreateFolderRequest) (resp *source.AddItemResponse, err error) {
	resp = source.NewAddItemResponse()
	resp.SetBase(base.NewBaseResp())
	// 生成唯一的标识
	key := util.UuidV4()
	folder := source.Item{
		Name:         req.Name,
		ItemType:     source.ItemType_FOLDER,
		Key:          key,
		Size:         0,
		ModifiedTime: time.Now().Unix(),
		Path:         req.Path,
	}
	success, err := s.SourceRepo.AddItem(req.SourceCategory, req.ParentId, &folder)
	if err != nil || success == false {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "创建文件夹错误"
		return
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "创建文件夹成功"
	resp.Item = &folder
	return
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
		return
	}
	if !success {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "添加失败"
		return
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "添加成功"
	return
}
