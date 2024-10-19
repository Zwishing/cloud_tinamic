package model

import (
	"cloud_tinamic/pkg"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	ServiceKey      string     `json:"service_key"`
	Title           string     `json:"title"`
	BBox            [4]float64 `json:"bbox"`
	Center          [2]float64 `json:"center"`
	Srid            int32      `json:"srid"`
	ServiceCategory int8       `json:"service_category"`
	Description     string     `json:"description"`
	Thumbnail       []byte     `json:"thumbnail"`
}

type Info struct {
	gorm.Model
	ServiceKey     string `json:"service_key"`
	Title          string `json:"title"`
	SourceKey      string `json:"sourceKey"`
	SourceSchema   string `json:"sourceSchema"`
	SourceCategory int    `json:"sourceCategory"`
	Srid           int32  `json:"srid"`
}

type Vector struct {
	gorm.Model
	SourceKey        string            `json:"service_key"`
	Title            string            `json:"title"`
	GeometryCategory string            `json:"geometry_category"`
	GeometryField    string            `json:"geometry_field"`
	FieldCount       int64             `json:"field_count"`
	RecordCount      int64             `json:"record_count"`
	Properties       map[string]string `json:"properties"`
}

func (c *Collection) Table() string {
	return pkg.ServiceCollectionTable
}

func (i *Info) Table() string {
	return pkg.ServiceInfoTable
}

func (i *Info) Vector() string {
	return pkg.ServiceVectorTable
}
