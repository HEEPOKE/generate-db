package repositories

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UtilitiesRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUtilitiesRepository(db *gorm.DB, rdb *redis.Client) *UtilitiesRepository {
	return &UtilitiesRepository{
		db:  db,
		rdb: rdb,
	}
}
