// Code generated by Kitex v0.11.3. DO NOT EDIT.

package servicecollection

import (
	collection "cloud_tinamic/kitex_gen/service/collection"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetCollections(ctx context.Context, pageSize int64, page int64, callOptions ...callopt.Option) (r *collection.GetCollectionsResponse, err error)
	AddCollection(ctx context.Context, sourceKey string, title string, callOptions ...callopt.Option) (r *collection.AddCollectionResponse, err error)
	Publish(ctx context.Context, req *collection.PublishRequest, callOptions ...callopt.Option) (err error)
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
	return &kServiceCollectionClient{
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

type kServiceCollectionClient struct {
	*kClient
}

func (p *kServiceCollectionClient) GetCollections(ctx context.Context, pageSize int64, page int64, callOptions ...callopt.Option) (r *collection.GetCollectionsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCollections(ctx, pageSize, page)
}

func (p *kServiceCollectionClient) AddCollection(ctx context.Context, sourceKey string, title string, callOptions ...callopt.Option) (r *collection.AddCollectionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddCollection(ctx, sourceKey, title)
}

func (p *kServiceCollectionClient) Publish(ctx context.Context, req *collection.PublishRequest, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Publish(ctx, req)
}
