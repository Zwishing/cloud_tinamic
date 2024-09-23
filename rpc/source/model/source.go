package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseInfo struct {
	gorm.Model
	Key            string `json:"key" column:"key"`
	Name           string `json:"name" column:"name"`
	SourceCategory int64  `json:"source_category" column:"source_category"`
}

func (b *BaseInfo) TableName() string {
	return "data_source.base_info"
}

type Storage struct {
	gorm.Model      `json:"gorm_._model"`
	Key             string    `json:"key" column:"key"`
	ParentKey       string    `json:"parent_key" column:"parent_key"`
	Name            string    `json:"name" column:"name"`
	StorageCategory int64     `json:"storage_category" column:"storage_category"`
	Size            int64     `json:"size" column:"size"`
	ModifiedTime    time.Time `json:"modified_time" column:"modified_time"`
	Path            string    `json:"path" column:"path"`
}

func (s *Storage) TableName() string {
	return "data_source.storage"
}

// AfterCreate 插入文件时，更新size
func (s *Storage) AfterCreate(tx *gorm.DB) (err error) {
	// 这里的1代表文件
	if s.StorageCategory == 1 && s.ParentKey != "" {
		parentKey := s.ParentKey
		sizeToAdd := s.Size

		// 使用 WITH RECURSIVE 递归查找父级文件夹并更新其大小
		err = tx.Exec(`
            WITH RECURSIVE parent_folders AS (
                SELECT key, parent_key
                FROM data_source.storage
                WHERE key = ?
                UNION ALL
                SELECT s.key, s.parent_key
                FROM data_source.storage s
                INNER JOIN parent_folders pf ON s.key = pf.parent_key
                WHERE s.storage_category = 2
            )
            UPDATE data_source.storage
            SET size = size + ?
            WHERE key IN (SELECT key FROM parent_folders);
        `, parentKey, sizeToAdd).Error

		if err != nil {
			return err
		}
	}
	return nil
}
