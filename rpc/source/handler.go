package main

import (
	"cloud_tinamic/kitex_gen/base"
	source "cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/rpc/source/pack"
	"cloud_tinamic/rpc/source/repo"
	"context"
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
	// TODO: Your code here...
	return
}

// PresignedUpload implements the SourceServiceImpl interface.
func (s *SourceServiceImpl) PresignedUpload(ctx context.Context, req *source.PresignedUploadResquest) (resp *source.PresignedUploadResponse, err error) {
	resp = source.NewPresignedUploadResponse()
	resp.SetBase(base.NewBaseResp())

	url, err := s.SourceRepo.PresignedUploadUrl(req.SourceType, req.Path, req.Name)
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
func (s *SourceServiceImpl) CreateFolder(ctx context.Context, req *source.CreateFolderRequest) (resp *source.CreateFolderResponse, err error) {
	// TODO: Your code here...
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

	success, err := s.SourceRepo.AddItem(req.SourceType, req.CurrentFolder, req.Item)
	if err != nil {
		resp.Base.Code = base.Code_SERVER_ERROR
		resp.Base.Msg = "添加错误"
		return nil, err
	}
	if !success {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "添加失败"
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "添加成功"
	return
}
