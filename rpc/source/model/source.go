package model

import (
	"cloud_tinamic/pkg"
	"gorm.io/gorm"
	"time"
)

type Info struct {
	gorm.Model
	Key            string `json:"key" column:"key"`
	Name           string `json:"name" column:"name"`
	SourceCategory int64  `json:"source_category" column:"source_category"`
}

func (b *Info) TableName() string {
	return pkg.SourceInfoTable
}

type Original struct {
	gorm.Model
	Key             string    `json:"key" column:"key"`
	ParentKey       string    `json:"parent_key" column:"parent_key"`
	Name            string    `json:"name" column:"name"`
	StorageCategory int64     `json:"storage_category" column:"storage_category"`
	Size            int64     `json:"size" column:"size"`
	ModifiedTime    time.Time `json:"modified_time" column:"modified_time"`
	Path            string    `json:"path" column:"path"`
}

func (s *Original) TableName() string {
	return pkg.SourceOriginalTable
}

type CloudOptimized struct {
	gorm.Model
	Key            string    `json:"key" column:"key"`
	SourceKey      string    `json:"sourceKey" column:"source_key"`
	SourceCategory int64     `json:"source_category" column:"source_category"`
	Size           int64     `json:"size" column:"size"`
	Path           string    `json:"path" column:"path"`
	ModifiedTime   time.Time `json:"modified_time" column:"modified_time"`
}

func (s *CloudOptimized) TableName() string {
	return pkg.SourceCloudOptimizedTable
}
