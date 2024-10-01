// Code generated by Kitex v0.11.0. DO NOT EDIT.

package vectorservice

import (
	vector "cloud_tinamic/kitex_gen/service/vector"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetCollections(ctx context.Context, pageSize int64, page int64, callOptions ...callopt.Option) (r *vector.GetCollectionsResponse, err error)
	Publish(ctx context.Context, req string, callOptions ...callopt.Option) (err error)
	GetTile(ctx context.Context, serviceKey string, x int32, y int32, zoom int8, ext string, params *vector.QueryParameters, callOptions ...callopt.Option) (r []byte, err error)
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
	return &kVectorServiceClient{
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

type kVectorServiceClient struct {
	*kClient
}

func (p *kVectorServiceClient) GetCollections(ctx context.Context, pageSize int64, page int64, callOptions ...callopt.Option) (r *vector.GetCollectionsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCollections(ctx, pageSize, page)
}

func (p *kVectorServiceClient) Publish(ctx context.Context, req string, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Publish(ctx, req)
}

func (p *kVectorServiceClient) GetTile(ctx context.Context, serviceKey string, x int32, y int32, zoom int8, ext string, params *vector.QueryParameters, callOptions ...callopt.Option) (r []byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTile(ctx, serviceKey, x, y, zoom, ext, params)
}
