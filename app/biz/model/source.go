package model

import (
	"cloud_tinamic/kitex_gen/data/source"
	"strings"
	"time"
)

type AddRequest struct {
}
type SourceItemResponse struct {
	ParentKey    string `json:"parentKey"`
	ItemType     string `json:"type"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Size         int64  `json:"size"`
	ModifiedTime string `json:"lastModified"`
}

type NewFolderRequest struct {
	SourceCategory string `json:"sourceCategory" validate:"required,oneof= vector imagery"`
	Key            string `json:"key" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Path           string `json:"path" validate:"required"`
}

type UploadRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// Items --------------------------------------------------------------
func Items(models []*source.Item) []*SourceItemResponse {
	items := make([]*SourceItemResponse, 0, len(models))
	for _, m := range models {
		if u := Item(m); u != nil {
			items = append(items, u)
		}
	}
	return items
}

func Item(model *source.Item) *SourceItemResponse {
	if model == nil {
		return nil
	}
	return &SourceItemResponse{
		ParentKey:    model.ParentKey,
		ItemType:     strings.ToLower(model.ItemType.String()),
		Key:          model.Key,
		Name:         model.Name,
		Path:         model.Path,
		Size:         model.Size,
		ModifiedTime: time.Unix(model.ModifiedTime, 0).Format("2006-01-02 15:04:05"),
	}
}
