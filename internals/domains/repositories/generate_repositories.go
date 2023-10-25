package repositories

import (
	"log"

	"github.com/HEEPOKE/generate-db/internals/core/cache"
	"github.com/HEEPOKE/generate-db/internals/core/utils"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"gorm.io/gorm"
)

type GenerateRepository struct {
	db *gorm.DB
}

func NewGenerateRepository(db *gorm.DB) *GenerateRepository {
	return &GenerateRepository{
		db: db,
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

func (g *GenerateRepository) GenerateData(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error) {
	quantity := int64(generateRequest.Quantity)
	results := utils.GenerateBatchData(int64(quantity), key, generateRequest)

	err := cache.SetKey(key, results)
	if err != nil {
		return nil, err
	}

	err = utils.CreateJSONFile(results, quantity, generateRequest.Table, key)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการสร้างไฟล์ JSON: %v", err)
		return nil, err
	}
	return results, nil
}
