package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseInfo struct {
	gorm.Model
	SourceId       string `json:"source_id" column:"source_id"`
	Name           string `json:"name" column:"name"`
	SourceCategory int64  `json:"source_category" column:"source_category"`
}

func (b *BaseInfo) TableName() string {
	return "data_source.base_info"
}

type Storage struct {
	gorm.Model      `json:"gorm_._model"`
	SourceId        string    `json:"source_id" column:"source_id"`
	ParentId        string    `json:"parent_id" column:"parent_id"`
	Name            string    `json:"name" column:"name"`
	StorageCategory int64     `json:"storage_category" column:"storage_category"`
	Key             string    `json:"key" column:"key"`
	Size            int64     `json:"size" column:"size"`
	ModifiedTime    time.Time `json:"modified_time" column:"modified_time"`
	Path            string    `json:"path" column:"path"`
}

func (s *Storage) TableName() string {
	return "data_source.storage"
}
