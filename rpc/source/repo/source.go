package repo

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg"
	"cloud_tinamic/pkg/storage"
	"cloud_tinamic/rpc/source/model"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ItemRepo interface {
	GetSiblingItemsByKey(key string) ([]*model.Original, error)
	GetChildrenItemsByKey(key string) ([]*model.Original, error)
	GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error)
	GetItemByName(key string, name string) ([]*model.Original, error)
	GetCountByName(key string, name string, itemType source.ItemType) (int64, error)
}

type StorageRepo interface {
	AddItem(SourceCategory source.SourceCategory, parentKey string, item *source.Item) (bool, error)
	DeleteItems(keys []string) (bool, error)
	GetPathByKey(key string) (string, error)
	GetHomeItemsBySourceCategory(sourceCategory source.SourceCategory) (string, []*model.Original, error)
	GetHomeKeyBySourceCategory(sourceCategory source.SourceCategory) (string, error)
}

type MinioRepo interface {
	PresignedUploadUrl(SourceCategory source.SourceCategory, path string, name string) (string, error)
	UploadToMinio(bucketName, path string, reader io.Reader, fileSize int64) error
}

type SourceRepo interface {
	ItemRepo
	StorageRepo
	CloudOptimizedRepo
	MinioRepo
}

type SourceRepoImpl struct {
	DB    *gorm.DB
	minio *storage.Storage
}

// NewSourceRepoImpl 创建一个新的 SourceRepoImpl 实例
func NewSourceRepoImpl(db *gorm.DB, minio *storage.Storage) SourceRepo {
	return &SourceRepoImpl{db, minio}
}

// UploadToMinio 上传文件到 Minio 存储
func (s *SourceRepoImpl) UploadToMinio(bucketName, path string, reader io.Reader, fileSize int64) error {
	_, err := s.minio.Upload(bucketName, path, reader, fileSize, minio.PutObjectOptions{})
	return err
}

// PresignedUploadUrl 获取预签名上传 URL
func (db *SourceRepoImpl) PresignedUploadUrl(SourceCategory source.SourceCategory, path string, name string) (string, error) {
	bucketName := strings.ToLower(SourceCategory.String())
	if db.minio.ObjExist(bucketName, name) {
		return "", fmt.Errorf("object already exists: %s/%s", bucketName, name)
	}

	presignedURL, err := db.minio.PutPresignedUrl(bucketName, path+name, 60)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL for %s/%s: %w", bucketName, path+name, err)
	}

	return presignedURL, nil
}

// GetChildrenItemsByKey 通过父键获取子项
func (s *SourceRepoImpl) GetChildrenItemsByKey(key string) ([]*model.Original, error) {
	var items []*model.Original
	err := s.DB.Table(pkg.SourceOriginalTable).Where("parent_key = ?", key).Find(&items).Error
	if err != nil {
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for parent_key: %s", key)
	}
	return items, nil
}

// GetSiblingItemsByKey 通过键获取同级项
func (s *SourceRepoImpl) GetSiblingItemsByKey(key string) ([]*model.Original, error) {
	var items []*model.Original
	err := s.DB.Where("parent_key = (?)", s.DB.Model(&model.Original{}).Select("parent_key").Where("key = ?", key)).Find(&items).Error
	if err != nil {
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for key: %s", key)
	}
	return items, nil
}

// GetItemByPath 通过路径获取项
func (db *SourceRepoImpl) GetItemByPath(sourceCategory source.SourceCategory, path string) ([]*source.Item, error) {
	return nil, nil
}

// AddItem 添加项
func (db *SourceRepoImpl) AddItem(sourceCategory source.SourceCategory, parentKey string, item *source.Item) (bool, error) {
	if parentKey == "" {
		var err error
		parentKey, err = db.GetHomeKeyBySourceCategory(sourceCategory)
		if err != nil {
			return false, fmt.Errorf("failed to get parent key: %w", err)
		}
	}

	baseInfo := model.Info{
		Key:            item.Key,
		Name:           item.Name,
		SourceCategory: int64(sourceCategory),
	}
	store := model.Original{
		Key:             item.Key,
		ParentKey:       parentKey,
		Name:            item.Name,
		StorageCategory: int64(item.ItemType),
		Size:            item.Size,
		ModifiedTime:    time.Unix(item.ModifiedTime, 0).Local(),
		Path:            item.Path,
	}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&baseInfo).Error; err != nil {
			return fmt.Errorf("failed to insert into base_info for uuid %s: %w", item.Key, err)
		}
		if err := tx.Create(&store).Error; err != nil {
			return fmt.Errorf("failed to insert into storage for uuid %s: %w", item.Key, err)
		}
		return nil
	})

	if err != nil {
		klog.Errorf("Transaction failed for uuid %s: %v", item.Key, err)
		return false, err
	}
	return true, nil
}

// GetHomeKeyBySourceCategory 通过源类别获取主键
func (db *SourceRepoImpl) GetHomeKeyBySourceCategory(sourceCategory source.SourceCategory) (string, error) {
	var parentKey string
	err := db.DB.Table(pkg.SourceOriginalTable+" AS s2").
		Joins("JOIN "+pkg.SourceInfoTable+" AS b ON s2.key = b.key").
		Select("s2.key").
		Where("s2.parent_key IS NULL AND b.source_category = ?", sourceCategory).
		Limit(1).Scan(&parentKey).Error
	if err != nil {
		return "", fmt.Errorf("failed to get parent key: %w", err)
	}
	if parentKey == "" {
		return "", fmt.Errorf("no root key found for source category: %v", sourceCategory)
	}
	return parentKey, nil
}

// DeleteItems 删除项
func (db *SourceRepoImpl) DeleteItems(keys []string) (bool, error) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		for _, table := range []interface{}{&model.Info{}, &model.Original{}} {
			result := tx.Where("key IN ?", keys).Delete(table)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("no records found for deletion in %T", table)
			}
		}
		return nil
	})

	if err != nil {
		klog.Errorf("Transaction failed for keys %v: %v", keys, err)
		return false, err
	}

	return true, nil
}

// GetPathByKey 通过键获取路径
func (s *SourceRepoImpl) GetPathByKey(key string) (string, error) {
	var path string
	err := s.DB.Table(pkg.SourceOriginalTable).Select("path").Where("key = ?", key).Scan(&path).Error
	return path, err
}

// GetItemByName 通过名称获取项
func (s *SourceRepoImpl) GetItemByName(parentKey string, name string) ([]*model.Original, error) {
	var items []*model.Original
	err := s.DB.Table(pkg.SourceOriginalTable).
		Where("parent_key = ? AND name = ?", parentKey, name).
		Find(&items).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query storage items: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for parent_key: %s and name: %s", parentKey, name)
	}
	return items, nil
}

// GetCountByName 通过名称获取计数
func (s *SourceRepoImpl) GetCountByName(key string, name string, itemType source.ItemType) (int64, error) {
	var count int64
	query := s.DB.Table(pkg.SourceOriginalTable).
		Where("storage_category = ? AND name = ?", itemType, name)

	if key == "" {
		query = query.Where("parent_key IS NULL")
	} else {
		query = query.Where("parent_key = ?", key)
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get count: %w", err)
	}

	return count, nil
}

// GetHomeItemsBySourceCategory 通过源类别获取主页项
func (s *SourceRepoImpl) GetHomeItemsBySourceCategory(sourceCategory source.SourceCategory) (string, []*model.Original, error) {
	var key string
	var storages []*model.Original

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		key, err = s.GetHomeKeyBySourceCategory(sourceCategory)
		if err != nil {
			return err
		}
		return tx.Table(pkg.SourceOriginalTable).
			Where("parent_key = ?", key).
			Find(&storages).Error
	})

	if err != nil {
		return "", nil, fmt.Errorf("failed to get top items by source category: %w", err)
	}

	if len(storages) == 0 {
		klog.Infof("No top items found for source category: %v", sourceCategory)
	}

	return key, storages, nil
}
