package model

import (
	"gorm.io/gorm"
)

type Collection struct {
	*gorm.Model
	ServiceKey      string     `json:"service_key"`
	Title           string     `json:"title"`
	BBox            [4]float64 `json:"bbox"`
	Center          [2]float64 `json:"center"`
	Srid            int32      `json:"srid"`
	ServiceCategory int8       `json:"service_category"`
	Description     string     `json:"description"`
	Thumbnail       []byte     `json:"thumbnail"`
}
