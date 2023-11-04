package repositories

import (
	"encoding/json"
	"log"

	"github.com/HEEPOKE/generate-db/internals/core/cache"
	"github.com/HEEPOKE/generate-db/internals/core/utils"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"gorm.io/gorm"
)

type GenerateRepository struct {
	db                  *gorm.DB
	utilitiesRepository *UtilitiesRepository
}

func NewGenerateRepository(db *gorm.DB, utilitiesRepository *UtilitiesRepository) *GenerateRepository {
	return &GenerateRepository{
		db:                  db,
		utilitiesRepository: utilitiesRepository,
	}
}

func (g *GenerateRepository) GetGenerateAll() ([]*models.Generate, error) {
	var generates []*models.Generate
	err := g.db.Find(&generates).Error
	if err != nil {
		return nil, err
	}
	return generates, nil
}

func (g *GenerateRepository) SaveDetailsGenerate(generate *models.Generate) error {
	return g.db.Create(generate).Error
}

func (g *GenerateRepository) GenerateData(data models.Generate, generateRequest *request.GenerateRequest) (models.JsonStructure, error) {
	results := utils.GenerateBatchData(data.Key, generateRequest)

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return models.JsonStructure{}, err
	}

	filePath, err := utils.CreateJSONFile(results, data)
	if err != nil {
		log.Printf("Error creating JSON file: %v", err)
		return models.JsonStructure{}, err
	}

	if err := g.utilitiesRepository.UpdatePathFileJson(data.Table, data.Key, filePath); err != nil {
		log.Printf("Error updating path and file JSON: %v", err)
		return models.JsonStructure{}, err
	}

	if err := cache.SetKey(data.Key, jsonData); err != nil {
		log.Printf("Error setting cache key: %v", err)
		return models.JsonStructure{}, err
	}

	return results, nil
}
