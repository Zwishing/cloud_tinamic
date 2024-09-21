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
	GetItemById(parentId string) ([]*model.Storage, error)
	GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error)
	AddItem(SourceCategory source.SourceCategory, dest string, item *source.Item) (bool, error)
	DeleteItem(SourceCategory source.SourceCategory, uuid string) (bool, error)
	GetPathById(sourceId string) (string, error)
	GetItemByName(parentId string, name string) ([]*model.Storage, error)
	GetCountByName(parentId string, name string, itemType source.ItemType) (int64, error)
	GetTopItemsBySourceCategory(sourceCateogry source.SourceCategory) ([]*model.Storage, error)
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

func (s *SourceRepoImpl) GetItemById(sourceId string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").
		Where("parent_id = ?", sourceId).
		Find(&items).Error
	if err != nil {
		// Changed from Fatalf to Errorf to avoid program termination
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	// Added a check for empty result
	if len(items) == 0 {
		klog.Infof("no items found for parent_id: %s", sourceId)
	}
	return items, nil
}

// GetItemByPath TODO
func (db *SourceRepoImpl) GetItemByPath(SourceCategory source.SourceCategory, path string) ([]*source.Item, error) {
	return nil, nil
}

func (db *SourceRepoImpl) AddItem(SourceCategory source.SourceCategory, parentId string, item *source.Item) (bool, error) {
	// Parse the modified time outside of the transaction
	modifiedTime := time.Unix(item.ModifiedTime, 0)

	// Open a transaction
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Prepare base_info data
		baseInfo := model.BaseInfo{
			SourceId:       item.Key,
			Name:           item.Name,
			SourceCategory: int64(SourceCategory),
		}

		// Prepare storage data
		store := model.Storage{
			SourceId:        item.Key,
			ParentId:        parentId,
			Name:            item.Name,
			StorageCategory: int64(item.ItemType),
			Key:             item.Key,
			Size:            item.Size,
			ModifiedTime:    modifiedTime,
			Path:            item.Path,
		}

		// Insert into base_info table
		if err := tx.Create(&baseInfo).Error; err != nil {
			return fmt.Errorf("failed to insert into base_info for uuid %s: %w", item.Key, err)
		}

		// Insert into storage table
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

func (db *SourceRepoImpl) DeleteItem(SourceCategory source.SourceCategory, uuid string) (bool, error) {
	// Use GORM's transaction processing
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Delete from both base_info and storage tables in a single query
		result := tx.Where("source_id = ?", uuid).Delete(&model.BaseInfo{}, &model.Storage{})
		if result.Error != nil {
			klog.Errorf("Failed to delete records for uuid %s: %v", uuid, result.Error)
			return result.Error
		}

		// Check if any records were affected
		if result.RowsAffected == 0 {
			klog.Warnf("No records found for deletion with uuid %s", uuid)
			return fmt.Errorf("no records found for deletion")
		}

		return nil
	})

	if err != nil {
		klog.Errorf("Transaction failed for uuid %s: %v", uuid, err)
		return false, err
	}

	return true, nil
}

func (s *SourceRepoImpl) GetPathById(sourceId string) (string, error) {
	var path string
	err := s.DB.Table("data_source.storage").
		Select("path").
		Where("source_id = ?", sourceId).
		Scan(&path).Error // 使用 Scan 以确保可以接收单个字段的值

	if err != nil {
		return "", err // 返回空字符串和错误
	}

	return path, nil // 返回查询到的路径
}

func (s *SourceRepoImpl) GetItemByName(parentId string, name string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").
		Where("parent_id = ? AND name = ?", parentId, name).
		Find(&items).Error
	if err != nil {
		klog.Errorf("failed to query storage items: %v", err)
		return nil, fmt.Errorf("failed to query storage items: %w", err)
	}
	if len(items) == 0 {
		klog.Infof("no items found for parent_id: %s and name: %s", parentId, name)
	}
	return items, nil
}

func (s *SourceRepoImpl) GetCountByName(parentId string, name string, itemType source.ItemType) (int64, error) {
	var count int64
	err := s.DB.Table("data_source.storage").
		Where("storage_category = ? AND parent_id = ? AND name = ?", itemType, parentId, name).
		Count(&count).Error
	if err != nil {
		err = fmt.Errorf("failed to get count: %w", err)
		klog.Error(err)
		return -1, err
	}
	return count, nil
}

func (s *SourceRepoImpl) GetTopItemsBySourceCategory(sourceCateogry source.SourceCategory) ([]*model.Storage, error) {
	var storages []*model.Storage
	err := s.DB.Table("data_source.storage AS s1").
		Joins("JOIN data_source.storage AS s2 ON s1.parent_id = s2.source_id").
		Joins("JOIN data_source.base_info AS b ON s2.source_id = b.source_id").
		Where("s2.parent_id IS NULL").
		Where("b.source_category = ?", sourceCateogry).
		Find(&storages).Error
	if err != nil {
		klog.Errorf("Failed to get top items by source category: %v", err)
		return nil, err
	}
	if len(storages) == 0 {
		klog.Infof("No top items found for source category: %v", sourceCateogry)
	}
	return storages, nil
}
