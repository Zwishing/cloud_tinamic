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

type DBRepo interface {
	GetSiblingItemsByKey(key string) ([]*model.Storage, error)
	GetChildrenItemsByKey(key string) ([]*model.Storage, error)
	GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error)
	AddItem(SourceCategory source.SourceCategory, parentKey string, item *source.Item) (bool, error)
	DeleteItems(key []string) (bool, error)
	GetPathByKey(key string) (string, error)
	GetItemByName(key string, name string) ([]*model.Storage, error)
	GetCountByName(key string, name string, itemType source.ItemType) (int64, error)
	GetHomeItemsBySourceCategory(sourceCategory source.SourceCategory) (string, []*model.Storage, error)
	GetHomeKeyBySourceCategory(sourceCategory source.SourceCategory) (string, error)
}

type MinioRepo interface {
	PresignedUploadUrl(SourceCategory source.SourceCategory, path string, name string) (string, error)
	UploadToMinio(bucketName, path, fileName string, reader io.Reader, fileSize int64) error
}

type SourceRepo interface {
	DBRepo
	MinioRepo
}

type SourceRepoImpl struct {
	DB    *gorm.DB
	minio *storage.Storage
}

// UploadToMinio implements SourceRepo.
func (s *SourceRepoImpl) UploadToMinio(bucketName, path, fileName string, reader io.Reader, fileSize int64) error {
	opts := minio.PutObjectOptions{}
	_, err := s.minio.Upload(bucketName, fileName, reader, fileSize, opts)
	return err
}

func NewSourceRepoImpl(db *gorm.DB, minio *storage.Storage) SourceRepo {
	return &SourceRepoImpl{
		db,
		minio,
	}
}

// GetChildrenItemsByKey 通过当前的key，查询所有属于当前的子类
func (s *SourceRepoImpl) GetChildrenItemsByKey(key string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").
		Where("parent_key = ? ", key).
		Find(&items).Error
	if err != nil {
		// Changed from Fatalf to Errorf to avoid program termination
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	// Added a check for empty result
	if len(items) == 0 {
		klog.Infof("no items found for parent_key: %s", key)
	}
	return items, nil
}

// GetSiblingItemsByKey 通过当前的key找到同一父节点的数据
func (s *SourceRepoImpl) GetSiblingItemsByKey(key string) ([]*model.Storage, error) {

	var items []*model.Storage
	err := s.DB.Where("parent_key = (?)",
		s.DB.Model(&model.Storage{}).Select("parent_key").Where("key = ?", key)).
		Find(&items).Error

	if err != nil {
		// Changed from Fatalf to Errorf to avoid program termination
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	// Added a check for empty result
	if len(items) == 0 {
		klog.Infof("no items found for parent_key: %s", key)
	}
	return items, nil

}

// GetItemByPath TODO
func (db *SourceRepoImpl) GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error) {
	return nil, nil
}

func (db *SourceRepoImpl) AddItem(SourceCategory source.SourceCategory, parentKey string, item *source.Item) (bool, error) {
	// 解析修改时间
	modifiedTime := time.Unix(item.ModifiedTime, 0)

	// 创建 baseInfo 和 store 对象
	baseInfo := model.BaseInfo{
		Key:            item.Key,
		Name:           item.Name,
		SourceCategory: int64(SourceCategory),
	}
	store := model.Storage{
		Key:             item.Key,
		ParentKey:       parentKey,
		Name:            item.Name,
		StorageCategory: int64(item.ItemType),
		Size:            item.Size,
		ModifiedTime:    modifiedTime,
		Path:            item.Path,
	}

	// 打开事务
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// 插入到 base_info 表
		if err := tx.Create(&baseInfo).Error; err != nil {
			klog.Errorf("failed to insert into base_info for uuid %s: %v", item.Key, err)
			return err
		}
		// 插入到 storage 表
		if err := tx.Create(&store).Error; err != nil {
			klog.Errorf("failed to insert into storage for uuid %s: %v", item.Key, err)
			return err
		}
		return nil
	})

	if err != nil {
		klog.Errorf("Transaction failed for uuid %s: %v", item.Key, err)
		return false, err
	}
	return true, nil
}
func (db *SourceRepoImpl) PresignedUploadUrl(SourceCategory source.SourceCategory, path string, name string) (string, error) {
	bucketName := strings.ToLower(SourceCategory.String())
	if db.minio.ObjExist(bucketName, name) {
		err := fmt.Errorf("object already exists: %s/%s", bucketName, name)
		klog.Errorf("failed to generate presigned URL: %v", err)
		return "", err
	}

	// Attempt to generate a presigned URL
	presignedURL, err := db.minio.PutPresignedUrl(bucketName, path+name, 60)
	if err != nil {
		err = fmt.Errorf("failed to generate presigned URL for %s/%s: %w", bucketName, path+name, err)
		klog.Error(err)
		return "", err
	}

	return presignedURL, nil
}

func (db *SourceRepoImpl) DeleteItems(keys []string) (bool, error) {
	// Use GORM's transaction processing
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Delete from both base_info and storage tables using IN clause
		result := tx.Where("key IN ?", keys).Delete(&model.BaseInfo{})
		if result.Error != nil {
			klog.Errorf("Failed to delete records for keys %v: %v", keys, result.Error)
			return result.Error
		}

		result = tx.Where("key IN ?", keys).Delete(&model.Storage{})
		if result.Error != nil {
			klog.Errorf("Failed to delete records from storage for keys %v: %v", keys, result.Error)
			return result.Error
		}

		// Check if any records were affected
		if result.RowsAffected == 0 {
			klog.Warnf("No records found for deletion with keys %v", keys)
			return fmt.Errorf("no records found for deletion")
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
	err := s.DB.Table("data_source.storage").
		Select("path").
		Where("key = ?", key).
		Scan(&path).Error // 使用 Scan 以确保可以接收单个字段的值

	if err != nil {
		return "", err // 返回空字符串和错误
	}

	return path, nil // 返回查询到的路径
}

func (s *SourceRepoImpl) GetItemByName(parentKey string, name string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").
		Where("parent_key = ? AND name = ?", parentKey, name).
		Find(&items).Error
	if err != nil {
		klog.Errorf("failed to query storage items: %v", err)
		return nil, fmt.Errorf("failed to query storage items: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for parent_key: %s and name: %s", parentKey, name)
	}
	return items, nil
}

func (s *SourceRepoImpl) GetCountByName(key string, name string, itemType source.ItemType) (int64, error) {
	var count int64
	err := s.DB.Table("data_source.storage").
		Where("storage_category = ? AND parent_key = ? AND name = ?", itemType, key, name).
		Count(&count).Error
	if err != nil {
		err = fmt.Errorf("failed to get count: %w", err)
		klog.Error(err)
		return -1, err
	}
	return count, nil
}

func (s *SourceRepoImpl) GetHomeItemsBySourceCategory(sourceCategory source.SourceCategory) (string, []*model.Storage, error) {
	var key string
	var storages []*model.Storage

	// 使用事务
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// 查询顶级键
		subQuery := tx.Table("data_source.storage AS s2").
			Joins("JOIN data_source.base_info AS b ON s2.key = b.key").
			Select("s2.key").
			Where("s2.parent_key IS NULL AND b.source_category = ?", sourceCategory).
			Limit(1)

		if err := subQuery.Scan(&key).Error; err != nil {
			return err
		}

		if key == "" {
			return fmt.Errorf("no top-level key found for source category: %v", sourceCategory)
		}

		// 查询子项
		return tx.Table("data_source.storage").
			Where("parent_key = ?", key).
			Find(&storages).Error
	})

	if err != nil {
		klog.Errorf("Failed to get top items by source category: %v", err)
		return "", nil, err
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
