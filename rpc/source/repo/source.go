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
		// Changed from Fatalf to Errorf to avoid program termination
		klog.Errorf("failed to query storages: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}
	// Added a check for empty result
	if len(items) == 0 {
		klog.Infof("no items found for parent_id: %s", uuid)
	}
	return items, nil
}

// GetItemByPath TODO
func (db *SourceRepoImpl) GetItemByPath(sourceType source.SourceType, path string) ([]*source.Item, error) {
	return nil, nil
}

func (db *SourceRepoImpl) AddItem(sourceType source.SourceType, dest string, item *source.Item) (bool, error) {
	// Parse the modified time outside of the transaction
	modifiedTime, err := time.Parse(time.RFC3339, item.ModifiedTime)
	if err != nil {
		klog.Errorf("Failed to parse time string '%s': %v", item.ModifiedTime, err)
		// Continue with zero time if parsing fails
		modifiedTime = time.Time{}
	}

	// Open a transaction
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		// Prepare base_info data
		baseInfo := model.BaseInfo{
			SourceId:       item.Key,
			Name:           item.Name,
			SourceCategory: int64(sourceType),
		}

		// Prepare storage data
		storage := model.Storage{
			SourceId:        item.Key,
			ParentId:        dest,
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
		if err := tx.Create(&storage).Error; err != nil {
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
