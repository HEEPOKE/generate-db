package repositories

import (
	"github.com/HEEPOKE/generate-db/internals/core/utils"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type GenerateRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewGenerateRepository(db *gorm.DB, rdb *redis.Client) *GenerateRepository {
	return &GenerateRepository{
		db:  db,
		rdb: rdb,
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

func (g *GenerateRepository) GenerateData(key string, generateRequest *request.GenerateRequest) (interface{}, error) {
	quantity := int64(generateRequest.Quantity)
	batchSize := int64(10000)

	numBatches := (quantity + batchSize - 1) / batchSize

	resultChan := make(chan []map[string]interface{}, numBatches)

	for i := int64(0); i < numBatches; i++ {
		batchSizeRemaining := quantity - i*batchSize
		batchSizeToGenerate := batchSizeRemaining
		if batchSizeRemaining > batchSize {
			batchSizeToGenerate = batchSize
		}

		go func(size int64) {
			results := utils.GenerateBatchData(size, generateRequest)
			resultChan <- results
		}(batchSizeToGenerate)
	}

	var results []map[string]interface{}
	for i := int64(0); i < numBatches; i++ {
		batchResults := <-resultChan
		results = append(results, batchResults...)
	}

	return results, nil
}
