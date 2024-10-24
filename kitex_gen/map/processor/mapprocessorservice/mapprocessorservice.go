// Code generated by Kitex v0.11.3. DO NOT EDIT.

package mapprocessorservice

import (
	processor "cloud_tinamic/kitex_gen/map/processor"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"VectorThumbnail": kitex.NewMethodInfo(
		vectorThumbnailHandler,
		newMapProcessorServiceVectorThumbnailArgs,
		newMapProcessorServiceVectorThumbnailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	mapProcessorServiceServiceInfo                = NewServiceInfo()
	mapProcessorServiceServiceInfoForClient       = NewServiceInfoForClient()
	mapProcessorServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return mapProcessorServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return mapProcessorServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return mapProcessorServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "MapProcessorService"
	handlerType := (*processor.MapProcessorService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "processor",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func vectorThumbnailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*processor.MapProcessorServiceVectorThumbnailArgs)
	realResult := result.(*processor.MapProcessorServiceVectorThumbnailResult)
	success, err := handler.(processor.MapProcessorService).VectorThumbnail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMapProcessorServiceVectorThumbnailArgs() interface{} {
	return processor.NewMapProcessorServiceVectorThumbnailArgs()
}

func newMapProcessorServiceVectorThumbnailResult() interface{} {
	return processor.NewMapProcessorServiceVectorThumbnailResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) VectorThumbnail(ctx context.Context, req *processor.VectorThumbnailRequest) (r *processor.VectorThumbnailRespose, err error) {
	var _args processor.MapProcessorServiceVectorThumbnailArgs
	_args.Req = req
	var _result processor.MapProcessorServiceVectorThumbnailResult
	if err = p.c.Call(ctx, "VectorThumbnail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
