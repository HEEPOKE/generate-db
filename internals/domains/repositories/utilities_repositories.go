package repositories

import (
	"encoding/json"

	"github.com/HEEPOKE/generate-db/internals/core/cache"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
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

func (u *UtilitiesRepository) CheckKeyData(key string) (models.JsonStructure, error) {
	data, err := cache.GetKey(key)
	if err != nil {
		return models.JsonStructure{}, err
	}

	jsonData := []byte(data)

	var result models.JsonStructure
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return models.JsonStructure{}, err
	}

	return result, nil
}
