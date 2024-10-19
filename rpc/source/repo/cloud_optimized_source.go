package repo

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/rpc/source/model"
	"fmt"
	"time"
)

// CloudOptimizedRepo 定义了云优化存储库的接口
type CloudOptimizedRepo interface {
	// GetCloudOptimizedSourcePathByKey 根据源键获取云优化的源路径
	GetCloudOptimizedSourcePathByKey(sourceKey string) ([]string, error)
	// AddCloudOptimizedItem 添加云优化项
	AddCloudOptimizedItem(sourceCategory source.SourceCategory, sourceKey string, unifiedKey string, path string, size int64) (bool, error)
}

// AddCloudOptimizedItem 添加云优化项到数据库
// 参数:
//   - sourceCategory: 源类别
//   - sourceKey: 源键
//   - unifiedKey: 统一键
//   - path: 路径
//   - size: 大小
// 返回:
//   - bool: 是否添加成功
//   - error: 错误信息
func (s *SourceRepoImpl) AddCloudOptimizedItem(sourceCategory source.SourceCategory, sourceKey string, unifiedKey string, path string, size int64) (bool, error) {
	unified := model.CloudOptimized{
		SourceKey:      sourceKey,
		Key:            unifiedKey,
		SourceCategory: int64(sourceCategory),
		Size:           size,
		Path:           path,
		ModifiedTime:   time.Now(),
	}

	if err := s.DB.Create(&unified).Error; err != nil {
		return false, fmt.Errorf("failed to add unified item: %w", err)
	}

	return true, nil
}

// GetCloudOptimizedSourcePathByKey 根据源键获取云优化的源路径
// 参数:
//   - sourceKey: 源键
// 返回:
//   - []string: 路径列表
//   - error: 错误信息
func (s *SourceRepoImpl) GetCloudOptimizedSourcePathByKey(sourceKey string) ([]string, error) {
	var paths []string
	err := s.DB.Model(&model.CloudOptimized{}).
		Select("path").
		Where("source_key = ?", sourceKey).
		Pluck("path", &paths).Error

	if err != nil {
		return nil, fmt.Errorf("failed to query unified paths: %w", err)
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("no paths found for source key: %s", sourceKey)
	}

	return paths, nil
}
