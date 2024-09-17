package handler

import (
	"cloud_tinamic/app/logger"
	"cloud_tinamic/kitex_gen/base/user/userservice"
	"cloud_tinamic/kitex_gen/data/source/sourceservice"
	"cloud_tinamic/kitex_gen/data/storage/storeservice"
	"github.com/cloudwego/kitex/client"
)

var (
	userClinet   userservice.Client
	geoClient    storeservice.Client
	sourceClient sourceservice.Client
)

func init() {
	// 创建 Kitex 客户端
	var err error
	userClinet, err = userservice.NewClient("base.user.userservice", client.WithHostPorts("0.0.0.0:8810"))
	geoClient, err = storeservice.NewClient("geo.data.storeservice", client.WithHostPorts("0.0.0.0:8089"))
	sourceClient, err = sourceservice.NewClient("data.source.sourceservice", client.WithHostPorts("0.0.0.0:8813"))
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("创建客户端失败")
	}
}
