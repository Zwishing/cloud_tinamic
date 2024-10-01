package repo

import (
	"gorm.io/gorm"
)

type ServiceRepo interface {
}
type ServiceRepoImpl struct {
	DB *gorm.DB
}

func NewServiceRepoImpl(db *gorm.DB) ServiceRepo {
	return &ServiceRepoImpl{
		DB: db,
	}
}

func (s *ServiceRepoImpl) GetCollections(page, pageSize int) {

}
