package repositories

import (
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

func (g *GenerateRepository) GenerateData(generateRequest *request.GenerateRequest) (interface{}, error) {
	var generateData request.GenerateRequest

	return nil, nil
}
