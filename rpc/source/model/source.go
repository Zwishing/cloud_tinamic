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
	gorm.Model
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

type Unified struct {
	gorm.Model
	Key            string    `json:"key" column:"key"`
	SourceKey      string    `json:"sourceKey" column:"source_key"`
	SourceCategory int64     `json:"source_category" column:"source_category"`
	Size           int64     `json:"size" column:"size"`
	Path           string    `json:"path" column:"path"`
	ModifiedTime   time.Time `json:"modified_time" column:"modified_time"`
}

func (s *Unified) TableName() string {
	return "data_source.unified"
}
