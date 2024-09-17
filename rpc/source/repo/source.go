package repo

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/storage"
	"cloud_tinamic/rpc/source/model"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"strings"
	"time"
)

type SourceRepo interface {
	GetItemById(uuid string) ([]*model.Storage, error)
	GetItemByPath(sourceType source.SourceType, path string) ([]*source.Item, error)
	AddItem(sourceType source.SourceType, dest string, item *source.Item) (bool, error)
	PresignedUploadUrl(sourceType source.SourceType, path string, name string) (string, error)
	DeleteItem(sourceType source.SourceType, uuid string) (bool, error)
}

type SourceRepoImpl struct {
	DB    *gorm.DB
	minio *storage.Storage
}

func NewSourceRepoImpl(db *gorm.DB, minio *storage.Storage) SourceRepo {
	return &SourceRepoImpl{
		db,
		minio,
	}
}

func (s *SourceRepoImpl) GetItemById(uuid string) ([]*model.Storage, error) {
	var items []*model.Storage
	err := s.DB.Table("data_source.storage").
		Where("parent_id = ?", uuid).
		Find(&items).Error
	if err != nil {
		klog.Fatalf("failed to query storages: %v", err)
		return nil, err
	}
	return items, nil
}

// GetItemByPath TODO
func (db *SourceRepoImpl) GetItemByPath(sourceType source.SourceType, path string) ([]*source.Item, error) {
	return nil, nil
}

func (db *SourceRepoImpl) AddItem(sourceType source.SourceType, dest string, item *source.Item) (bool, error) {
	// Open a transaction
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Insert into base_info table
		if err := tx.Create(&model.BaseInfo{
			SourceId:       item.Key,
			Name:           item.Name,
			SourceCategory: int64(sourceType),
		}).Error; err != nil {
			klog.Errorf("Failed to insert into base_info for uuid %s: %v", item.Key, err)
			return fmt.Errorf("failed to insert into base_info for uuid %s: %w", item.Key, err)
		}
		modifiedTime, err := time.Parse(time.RFC3339, item.ModifiedTime)
		if err != nil {
			klog.Errorf("Failed to parse time string '%s': %v", item.ModifiedTime, err)
		}

		// Insert into storage table
		if err := tx.Create(&model.Storage{
			SourceId:        item.Key,
			ParentId:        dest,
			Name:            item.Name,
			StorageCategory: int64(item.ItemType),
			Key:             item.Key,
			Size:            item.Size,
			ModifiedTime:    modifiedTime,
			Path:            item.Path,
		}).Error; err != nil {
			klog.Errorf("Failed to insert into storage for uuid %s: %v", item.Key, err)
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

func (db *SourceRepoImpl) PresignedUploadUrl(sourceType source.SourceType, path string, name string) (string, error) {
	bucketName := strings.ToLower(sourceType.String())
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

func (db *SourceRepoImpl) DeleteItem(sourceType source.SourceType, uuid string) (bool, error) {
	// 使用 GORM 的事务处理
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// 从 base_info 表中删除
		if err := tx.Where("source_id = ?", uuid).Delete(&model.BaseInfo{}).Error; err != nil {
			klog.Errorf("Failed to delete from base_info table for uuid %s: %v", uuid, err)
			return err
		}

		// 从 storage 表中删除
		if err := tx.Where("source_id = ?", uuid).Delete(&model.Storage{}).Error; err != nil {
			klog.Errorf("Failed to delete from storage table for uuid %s: %v", uuid, err)
			return err
		}

		return nil
	})

	if err != nil {
		klog.Errorf("Transaction failed for uuid %s: %v", uuid, err)
		return false, err
	}

	return true, nil
}
