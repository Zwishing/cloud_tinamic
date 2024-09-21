package pack

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/rpc/source/model"
)

func Storages(models []*model.Storage) []*source.Item {
	storages := make([]*source.Item, 0, len(models))
	for _, m := range models {
		if u := Storage(m); u != nil {
			storages = append(storages, u)
		}
	}
	return storages
}

func Storage(m *model.Storage) *source.Item {
	if m == nil {
		return nil
	}
	return &source.Item{
		ParentKey:    m.ParentKey,
		Name:         m.Name,
		ItemType:     source.ItemType(m.StorageCategory),
		Key:          m.Key,
		Size:         m.Size,
		ModifiedTime: m.ModifiedTime.Unix(),
		Path:         m.Path,
	}
}
