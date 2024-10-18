package repo

import (
	"cloud_tinamic/kitex_gen/data/source"
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
	GetSiblingItemsByKey(key string) ([]*model.Storage, error)
	GetChildrenItemsByKey(key string) ([]*model.Storage, error)
	GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error)
	GetItemByName(key string, name string) ([]*model.Storage, error)
	GetCountByName(key string, name string, itemType source.ItemType) (int64, error)
}

type StorageRepo interface {
	AddItem(SourceCategory source.SourceCategory, parentKey string, item *source.Item) (bool, error)
	DeleteItems(keys []string) (bool, error)
	GetPathByKey(key string) (string, error)
	GetHomeItemsBySourceCategory(sourceCategory source.SourceCategory) (string, []*model.Storage, error)
	GetHomeKeyBySourceCategory(sourceCategory source.SourceCategory) (string, error)
}

type UnifiedRepo interface {
	GetUnifiedSourcePathByKey(sourceKey string) ([]string, error)
	AddCloudOptimizedItem(sourceCategory source.SourceCategory, sourceKey string, unifiedKey string, path string, size int64) (bool, error)
}

type MinioRepo interface {
	PresignedUploadUrl(SourceCategory source.SourceCategory, path string, name string) (string, error)
	UploadToMinio(bucketName, path string, reader io.Reader, fileSize int64) error
}

type SourceRepo interface {
	ItemRepo
	StorageRepo
	UnifiedRepo
	MinioRepo
}

type SourceRepoImpl struct {
	DB    *gorm.DB
	minio *storage.Storage
}

func NewSourceRepoImpl(db *gorm.DB, minio *storage.Storage) SourceRepo {
	return &SourceRepoImpl{db, minio}
}

func (s *SourceRepoImpl) UploadToMinio(bucketName, path string, reader io.Reader, fileSize int64) error {
	_, err := s.minio.Upload(bucketName, path, reader, fileSize, minio.PutObjectOptions{})
	return err
}

func (s *SourceRepoImpl) GetChildrenItemsByKey(key string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").Where("parent_key = ?", key).Find(&items).Error
	if err != nil {
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for parent_key: %s", key)
	}
	return items, nil
}

func (s *SourceRepoImpl) GetSiblingItemsByKey(key string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Where("parent_key = (?)", s.DB.Model(&model.Storage{}).Select("parent_key").Where("key = ?", key)).Find(&items).Error
	if err != nil {
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for key: %s", key)
	}
	return items, nil
}

func (db *SourceRepoImpl) GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error) {
	return nil, nil
}

func (db *SourceRepoImpl) AddItem(sourceCategory source.SourceCategory, parentKey string, item *source.Item) (bool, error) {
	if parentKey == "" {
		var err error
		parentKey, err = db.getSourceCategoryRootKey(sourceCategory)
		if err != nil {
			return false, fmt.Errorf("failed to get parent key: %w", err)
		}
	}

	baseInfo := model.BaseInfo{
		Key:            item.Key,
		Name:           item.Name,
		SourceCategory: int64(sourceCategory),
	}
	store := model.Storage{
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

func (db *SourceRepoImpl) getSourceCategoryRootKey(sourceCategory source.SourceCategory) (string, error) {
	var parentKey string
	err := db.DB.Table("data_source.storage AS s2").
		Joins("JOIN data_source.base_info AS b ON s2.key = b.key").
		Select("s2.key").
		Where("s2.parent_key IS NULL AND b.source_category = ?", sourceCategory).
		Limit(1).Scan(&parentKey).Error
	if err != nil {
		return "", fmt.Errorf("failed to get parent key: %w", err)
	}
	return parentKey, nil
}

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

func (db *SourceRepoImpl) DeleteItems(keys []string) (bool, error) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		for _, table := range []interface{}{&model.BaseInfo{}, &model.Storage{}} {
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

func (s *SourceRepoImpl) GetPathByKey(key string) (string, error) {
	var path string
	err := s.DB.Table("data_source.storage").Select("path").Where("key = ?", key).Scan(&path).Error
	return path, err
}

func (s *SourceRepoImpl) GetUnifiedSourcePathByKey(sourceKey string) ([]string, error) {
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

func (s *SourceRepoImpl) GetItemByName(parentKey string, name string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").
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

func (s *SourceRepoImpl) GetCountByName(key string, name string, itemType source.ItemType) (int64, error) {
	var count int64
	query := s.DB.Table("data_source.storage").
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

func (s *SourceRepoImpl) GetHomeItemsBySourceCategory(sourceCategory source.SourceCategory) (string, []*model.Storage, error) {
	var key string
	var storages []*model.Storage

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("data_source.storage AS s2").
			Joins("JOIN data_source.base_info AS b ON s2.key = b.key").
			Select("s2.key").
			Where("s2.parent_key IS NULL AND b.source_category = ?", sourceCategory).
			Limit(1).Scan(&key).Error; err != nil {
			return err
		}

		if key == "" {
			return fmt.Errorf("no top-level key found for source category: %v", sourceCategory)
		}

		return tx.Table("data_source.storage").
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

func (s *SourceRepoImpl) GetHomeKeyBySourceCategory(sourceCategory source.SourceCategory) (string, error) {
	var key string

	err := s.DB.Model(&model.BaseInfo{}).
		Select("s.key").
		Joins("JOIN data_source.storage AS s ON s.key = base_info.key").
		Where("s.parent_key IS NULL AND base_info.source_category = ?", sourceCategory).
		Limit(1).
		Scan(&key).Error

	if err != nil {
		return "", err
	}

	if key == "" {
		return "", fmt.Errorf("no key found for the given source category")
	}

	return key, nil
}

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
