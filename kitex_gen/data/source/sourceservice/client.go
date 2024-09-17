// Code generated by Kitex v0.11.0. DO NOT EDIT.

package sourceservice

import (
	source "cloud_tinamic/kitex_gen/data/source"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Upload(ctx context.Context, req *source.UploadRequest, callOptions ...callopt.Option) (r *source.UploadResponse, err error)
	PresignedUpload(ctx context.Context, req *source.PresignedUploadResquest, callOptions ...callopt.Option) (r *source.PresignedUploadResponse, err error)
	GetItem(ctx context.Context, req *source.GetItemRequest, callOptions ...callopt.Option) (r *source.GetItemResponse, err error)
	CreateFolder(ctx context.Context, req *source.CreateFolderRequest, callOptions ...callopt.Option) (r *source.CreateFolderResponse, err error)
	DeleteItem(ctx context.Context, req *source.DeleteItemRequest, callOptions ...callopt.Option) (r *source.DeleteItemResponse, err error)
	AddItem(ctx context.Context, req *source.AddItemRequest, callOptions ...callopt.Option) (r *source.AddItemResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kSourceServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSourceServiceClient struct {
	*kClient
}

func (p *kSourceServiceClient) Upload(ctx context.Context, req *source.UploadRequest, callOptions ...callopt.Option) (r *source.UploadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Upload(ctx, req)
}

func (p *kSourceServiceClient) PresignedUpload(ctx context.Context, req *source.PresignedUploadResquest, callOptions ...callopt.Option) (r *source.PresignedUploadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PresignedUpload(ctx, req)
}

func (p *kSourceServiceClient) GetItem(ctx context.Context, req *source.GetItemRequest, callOptions ...callopt.Option) (r *source.GetItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetItem(ctx, req)
}

func (p *kSourceServiceClient) CreateFolder(ctx context.Context, req *source.CreateFolderRequest, callOptions ...callopt.Option) (r *source.CreateFolderResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFolder(ctx, req)
}

func (p *kSourceServiceClient) DeleteItem(ctx context.Context, req *source.DeleteItemRequest, callOptions ...callopt.Option) (r *source.DeleteItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteItem(ctx, req)
}

func (p *kSourceServiceClient) AddItem(ctx context.Context, req *source.AddItemRequest, callOptions ...callopt.Option) (r *source.AddItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddItem(ctx, req)
}