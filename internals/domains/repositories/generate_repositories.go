package repositories

import (
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
