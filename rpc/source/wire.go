//go:build wireinject
// +build wireinject

// wire.go
package main

import (
	"cloud_tinamic/rpc/source/repo"
	"github.com/google/wire"
)

// 定义数据库的 ProviderSet
var DatabaseProviders = wire.NewSet(
	repo.NewDB,             // 共享数据库连接池
	repo.NewSourceRepoImpl, // 依赖数据库的存储库
)

// 定义 MinIO 的 ProviderSet
var MinioProviders = wire.NewSet(
	repo.NewMinio, // 共享 MinIO 实例
)

// 组合所有 Providers，复用依赖
var DataProviders = wire.NewSet(
	DatabaseProviders,
	MinioProviders,
)

// SourceService 注入优化
func InitSourceService() (*SourceServiceImpl, error) {
	wire.Build(
		DataProviders, // 复用数据库和 MinIO 的依赖
		NewGeoServiceClient,
		NewSourceServiceImpl, // 注入服务
	)
	return nil, nil
}

// NSQ Client 注入优化
func InitNsqClient() *repo.NsqConsumer {
	wire.Build(
		DataProviders,       // 只注入数据库相关依赖
		repo.NewNsqConfig,   // NSQ 配置
		repo.NewNSQHandler,  // 初始化 NSQ 处理器
		repo.NewNsqConsumer, // 初始化 NSQ 消费者
	)
	return nil
}
