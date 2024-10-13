// Code generated by Kitex v0.11.3. DO NOT EDIT.

package sourceservice

import (
	source "cloud_tinamic/kitex_gen/data/source"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetNextItems": kitex.NewMethodInfo(
		getNextItemsHandler,
		newSourceServiceGetNextItemsArgs,
		newSourceServiceGetNextItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPreviousItems": kitex.NewMethodInfo(
		getPreviousItemsHandler,
		newSourceServiceGetPreviousItemsArgs,
		newSourceServiceGetPreviousItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetHomeItems": kitex.NewMethodInfo(
		getHomeItemsHandler,
		newSourceServiceGetHomeItemsArgs,
		newSourceServiceGetHomeItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteItems": kitex.NewMethodInfo(
		deleteItemsHandler,
		newSourceServiceDeleteItemsArgs,
		newSourceServiceDeleteItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AddItem": kitex.NewMethodInfo(
		addItemHandler,
		newSourceServiceAddItemArgs,
		newSourceServiceAddItemResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateFolder": kitex.NewMethodInfo(
		createFolderHandler,
		newSourceServiceCreateFolderArgs,
		newSourceServiceCreateFolderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Upload": kitex.NewMethodInfo(
		uploadHandler,
		newSourceServiceUploadArgs,
		newSourceServiceUploadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"PresignedUpload": kitex.NewMethodInfo(
		presignedUploadHandler,
		newSourceServicePresignedUploadArgs,
		newSourceServicePresignedUploadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetSourcePath": kitex.NewMethodInfo(
		getSourcePathHandler,
		newSourceServiceGetSourcePathArgs,
		newSourceServiceGetSourcePathResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	sourceServiceServiceInfo                = NewServiceInfo()
	sourceServiceServiceInfoForClient       = NewServiceInfoForClient()
	sourceServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return sourceServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return sourceServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return sourceServiceServiceInfoForClient
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
	serviceName := "SourceService"
	handlerType := (*source.SourceService)(nil)
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
		"PackageName": "source",
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

func getNextItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceGetNextItemsArgs)
	realResult := result.(*source.SourceServiceGetNextItemsResult)
	success, err := handler.(source.SourceService).GetNextItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceGetNextItemsArgs() interface{} {
	return source.NewSourceServiceGetNextItemsArgs()
}

func newSourceServiceGetNextItemsResult() interface{} {
	return source.NewSourceServiceGetNextItemsResult()
}

func getPreviousItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceGetPreviousItemsArgs)
	realResult := result.(*source.SourceServiceGetPreviousItemsResult)
	success, err := handler.(source.SourceService).GetPreviousItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceGetPreviousItemsArgs() interface{} {
	return source.NewSourceServiceGetPreviousItemsArgs()
}

func newSourceServiceGetPreviousItemsResult() interface{} {
	return source.NewSourceServiceGetPreviousItemsResult()
}

func getHomeItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceGetHomeItemsArgs)
	realResult := result.(*source.SourceServiceGetHomeItemsResult)
	success, err := handler.(source.SourceService).GetHomeItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceGetHomeItemsArgs() interface{} {
	return source.NewSourceServiceGetHomeItemsArgs()
}

func newSourceServiceGetHomeItemsResult() interface{} {
	return source.NewSourceServiceGetHomeItemsResult()
}

func deleteItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceDeleteItemsArgs)
	realResult := result.(*source.SourceServiceDeleteItemsResult)
	success, err := handler.(source.SourceService).DeleteItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceDeleteItemsArgs() interface{} {
	return source.NewSourceServiceDeleteItemsArgs()
}

func newSourceServiceDeleteItemsResult() interface{} {
	return source.NewSourceServiceDeleteItemsResult()
}

func addItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceAddItemArgs)
	realResult := result.(*source.SourceServiceAddItemResult)
	success, err := handler.(source.SourceService).AddItem(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceAddItemArgs() interface{} {
	return source.NewSourceServiceAddItemArgs()
}

func newSourceServiceAddItemResult() interface{} {
	return source.NewSourceServiceAddItemResult()
}

func createFolderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceCreateFolderArgs)
	realResult := result.(*source.SourceServiceCreateFolderResult)
	success, err := handler.(source.SourceService).CreateFolder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceCreateFolderArgs() interface{} {
	return source.NewSourceServiceCreateFolderArgs()
}

func newSourceServiceCreateFolderResult() interface{} {
	return source.NewSourceServiceCreateFolderResult()
}

func uploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceUploadArgs)
	realResult := result.(*source.SourceServiceUploadResult)
	success, err := handler.(source.SourceService).Upload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServiceUploadArgs() interface{} {
	return source.NewSourceServiceUploadArgs()
}

func newSourceServiceUploadResult() interface{} {
	return source.NewSourceServiceUploadResult()
}

func presignedUploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServicePresignedUploadArgs)
	realResult := result.(*source.SourceServicePresignedUploadResult)
	success, err := handler.(source.SourceService).PresignedUpload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSourceServicePresignedUploadArgs() interface{} {
	return source.NewSourceServicePresignedUploadArgs()
}

func newSourceServicePresignedUploadResult() interface{} {
	return source.NewSourceServicePresignedUploadResult()
}

func getSourcePathHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*source.SourceServiceGetSourcePathArgs)
	realResult := result.(*source.SourceServiceGetSourcePathResult)
	success, err := handler.(source.SourceService).GetSourcePath(ctx, realArg.Key)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newSourceServiceGetSourcePathArgs() interface{} {
	return source.NewSourceServiceGetSourcePathArgs()
}

func newSourceServiceGetSourcePathResult() interface{} {
	return source.NewSourceServiceGetSourcePathResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetNextItems(ctx context.Context, req *source.GetItemsRequest) (r *source.GetItemsResponse, err error) {
	var _args source.SourceServiceGetNextItemsArgs
	_args.Req = req
	var _result source.SourceServiceGetNextItemsResult
	if err = p.c.Call(ctx, "GetNextItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPreviousItems(ctx context.Context, req *source.GetItemsRequest) (r *source.GetItemsResponse, err error) {
	var _args source.SourceServiceGetPreviousItemsArgs
	_args.Req = req
	var _result source.SourceServiceGetPreviousItemsResult
	if err = p.c.Call(ctx, "GetPreviousItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetHomeItems(ctx context.Context, req *source.GetHomeItemsRequest) (r *source.GetItemsResponse, err error) {
	var _args source.SourceServiceGetHomeItemsArgs
	_args.Req = req
	var _result source.SourceServiceGetHomeItemsResult
	if err = p.c.Call(ctx, "GetHomeItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteItems(ctx context.Context, req *source.DeleteItemsRequest) (r *source.DeleteItemsResponse, err error) {
	var _args source.SourceServiceDeleteItemsArgs
	_args.Req = req
	var _result source.SourceServiceDeleteItemsResult
	if err = p.c.Call(ctx, "DeleteItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddItem(ctx context.Context, req *source.AddItemRequest) (r *source.AddItemResponse, err error) {
	var _args source.SourceServiceAddItemArgs
	_args.Req = req
	var _result source.SourceServiceAddItemResult
	if err = p.c.Call(ctx, "AddItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateFolder(ctx context.Context, req *source.CreateFolderRequest) (r *source.AddItemResponse, err error) {
	var _args source.SourceServiceCreateFolderArgs
	_args.Req = req
	var _result source.SourceServiceCreateFolderResult
	if err = p.c.Call(ctx, "CreateFolder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Upload(ctx context.Context, req *source.UploadRequest) (r *source.UploadResponse, err error) {
	var _args source.SourceServiceUploadArgs
	_args.Req = req
	var _result source.SourceServiceUploadResult
	if err = p.c.Call(ctx, "Upload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PresignedUpload(ctx context.Context, req *source.PresignedUploadResquest) (r *source.PresignedUploadResponse, err error) {
	var _args source.SourceServicePresignedUploadArgs
	_args.Req = req
	var _result source.SourceServicePresignedUploadResult
	if err = p.c.Call(ctx, "PresignedUpload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSourcePath(ctx context.Context, key string) (r string, err error) {
	var _args source.SourceServiceGetSourcePathArgs
	_args.Key = key
	var _result source.SourceServiceGetSourcePathResult
	if err = p.c.Call(ctx, "GetSourcePath", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
