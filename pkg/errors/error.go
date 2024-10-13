package errors

import (
	"fmt"
	"strings"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ErrorCode = int32

// 定义错误类型
const (
	InvalidInputCode       ErrorCode = 4001
	InvalidParametersCode  ErrorCode = 4002
	InternalServerCode     ErrorCode = 5001
	NetworkIssueCode       ErrorCode = 6001
	NotFoundCode           ErrorCode = 7001
	UnauthorizedCode       ErrorCode = 7002
	OutOfMemoryCode        ErrorCode = 8001
	ServiceUnavailableCode ErrorCode = 9001
	TimeoutCode            ErrorCode = 10001
	QueryFailedCode        ErrorCode = 11001
	DatabaseErrorCode      ErrorCode = 11002
	ResourceExistsCode     ErrorCode = 12001
	OperationFailedCode    ErrorCode = 13001
	UnknownCode            ErrorCode = 14001
)

// 定义常量错误信息
var errorMessages = map[ErrorCode]string{
	InvalidInputCode:       "无效的输入参数",
	InternalServerCode:     "内部服务器错误",
	NetworkIssueCode:       "网络连接问题",
	NotFoundCode:           "资源未找到",
	UnauthorizedCode:       "未经授权的访问",
	OutOfMemoryCode:        "内存不足",
	ServiceUnavailableCode: "第三方服务不可用",
	TimeoutCode:            "请求超时",
	QueryFailedCode:        "查询失败",
	ResourceExistsCode:     "资源已存在",
	OperationFailedCode:    "操作失败",
}

// Kerror 根据错误代码和消息返回相应的错误
func Kerror(code ErrorCode, msg string) error {
	errMsg, exists := errorMessages[code]
	if !exists {
		return kerrors.NewBizStatusError(UnknownCode, "未知错误")
	}
	var builder strings.Builder
	builder.WriteString(errMsg)
	builder.WriteString(": ")
	builder.WriteString(msg)
	return kerrors.NewBizStatusError(code, builder.String())
}

func Kerrorf(code ErrorCode, msg string, args ...any) error {
	errMsg, exists := errorMessages[code]
	if !exists {
		return kerrors.NewBizStatusError(UnknownCode, "未知错误")
	}

	var builder strings.Builder
	builder.WriteString(errMsg)
	builder.WriteString(": ")
	builder.WriteString(fmt.Sprintf(msg, args...))

	return kerrors.NewBizStatusError(code, builder.String())
}

// 预定义的错误常量
var (
	ErrInvalidInput       = Kerror(InvalidInputCode, "无效的输入参数")
	ErrInternalServer     = Kerror(InternalServerCode, "内部服务器错误")
	ErrNetworkIssue       = Kerror(NetworkIssueCode, "网络连接问题")
	ErrNotFound           = Kerror(NotFoundCode, "资源未找到")
	ErrUnauthorized       = Kerror(UnauthorizedCode, "未经授权的访问")
	ErrOutOfMemory        = Kerror(OutOfMemoryCode, "内存不足")
	ErrServiceUnavailable = Kerror(ServiceUnavailableCode, "第三方服务不可用")
	ErrTimeout            = Kerror(TimeoutCode, "请求超时")
	ErrQueryFailed        = Kerror(QueryFailedCode, "查询失败")
	ErrResourceExists     = Kerror(ResourceExistsCode, "资源已存在")
	ErrOperationFailed    = Kerror(OperationFailedCode, "操作失败")
)
