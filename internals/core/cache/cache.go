package cache

import (
	"context"
	"time"

	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/pkg/databases"
)

func SetKey(key string, value models.JsonStructure) error {
	client := databases.ConnectToRedis()
	expiration := 24 * time.Hour
	return client.Set(context.Background(), key, value, expiration).Err()
}

func GetKey(key string) (string, error) {
	client := databases.ConnectToRedis()
	return client.Get(context.Background(), key).Result()
}
