package main

import (
	vector "cloud_tinamic/kitex_gen/service/vector"
	. "cloud_tinamic/pkg/errors"
	"cloud_tinamic/rpc/vector_service/repo"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"time"
)

// VectorServiceImpl implements the last service interface defined in the IDL.
type VectorServiceImpl struct {
}

// Publish implements the VectorServiceImpl interface.
func (s *VectorServiceImpl) Publish(ctx context.Context, req string) (err error) {
	// TODO: Your code here...
	return
}

// GetTile implements the VectorServiceImpl interface.
func (s *VectorServiceImpl) GetTile(ctx context.Context, serviceKey string,
	x, y int32, z int8, ext string, params *vector.QueryParameters) (resp []byte, err error) {
	// 使用go-cache本地缓存
	lyr, err := repo.GetLayer(serviceKey)
	if err != nil {
		// 从数据库查询获取
		return nil, Kerror(NotFoundCode, "")
	}

	tile, errTile := repo.MakeTile(x, y, z, ext)
	if errTile != nil {
		return nil, errTile
	}
	klog.Tracef("requestTile: %s", tile.String())
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("DbTimeout")*time.Second)
	defer cancel()

	tileRequest := lyr.GetTileRequest(tile, params)
	mvt, errMvt := repo.DBTileRequest(ctx, &tileRequest)
	if errMvt != nil {
		return nil, errMvt
	}

	return mvt, nil
}

func (s *VectorServiceImpl) GetCollections(ctx context.Context, pageSize int64, page int64) (resp *vector.GetCollectionsResponse, err error) {
	// TODO: Your code here...
	return
}
