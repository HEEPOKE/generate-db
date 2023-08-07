package repositories_test

import (
	"testing"

	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/HEEPOKE/generate-db/internals/domains/repositories"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newMockGenerateRepository() *repositories.GenerateRepository {
	db := &gorm.DB{}
	rdb := &redis.Client{}
	return repositories.NewGenerateRepository(db, rdb)
}

func TestGetGenerateAll(t *testing.T) {
	mockRepo := newMockGenerateRepository()

	testGenerates := []*models.Generate{
		{Key: "key1", Table: "table1", Quantity: 10},
		{Key: "key2", Table: "table2", Quantity: 20},
	}

	for _, generate := range testGenerates {
		mockRepo.SaveDetailsGenerate(generate)
	}

	generates, err := mockRepo.GetGenerateAll()

	assert.NoError(t, err)
	assert.NotNil(t, generates)
	assert.Len(t, generates, len(testGenerates))
}

func TestSaveDetailsGenerate(t *testing.T) {
	mockRepo := newMockGenerateRepository()

	generate := &models.Generate{
		Key:      "test_key",
		Table:    "test_table",
		Quantity: 50,
	}

	err := mockRepo.SaveDetailsGenerate(generate)

	assert.NoError(t, err)
}

func TestGenerateData(t *testing.T) {
	mockRepo := newMockGenerateRepository()

	generateRequest := &request.GenerateRequest{
		Quantity: 100,
	}

	key := "test_key"

	results, err := mockRepo.GenerateData(key, generateRequest)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results) <= 100)

	generateRequest2 := &request.GenerateRequest{
		Quantity: 1000,
	}
	key2 := "test_key2"

	results2, err2 := mockRepo.GenerateData(key2, generateRequest2)

	assert.Error(t, err2)
	assert.Nil(t, results2)
}
