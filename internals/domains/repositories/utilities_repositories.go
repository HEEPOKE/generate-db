package repositories

import (
	"gorm.io/gorm"
)

type UtilitiesRepository struct {
	db *gorm.DB
}

func NewUtilitiesRepository(db *gorm.DB) *UtilitiesRepository {
	return &UtilitiesRepository{
		db: db,
	}
}
