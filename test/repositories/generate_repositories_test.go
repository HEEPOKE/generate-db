package repositories_test

import (
	"testing"

	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/HEEPOKE/generate-db/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetGenerateAll(t *testing.T) {
	mockRepo := &mocks.MockGenerateRepository{}

	mockRepo.GetGenerateAllFunc = func() ([]*models.Generate, error) {
		return []*models.Generate{
			{ID: 1, Key: "key1", Table: "table1", Quantity: 10},
			{ID: 2, Key: "key2", Table: "table2", Quantity: 20},
		}, nil
	}

	generates, err := mockRepo.GetGenerateAll()

	assert.NoError(t, err)
	assert.NotNil(t, generates)
	assert.Len(t, generates, 2)
}

func TestSaveDetailsGenerate(t *testing.T) {
	mockRepo := &mocks.MockGenerateRepository{}

	mockRepo.SaveDetailsGenerateFunc = func(generate *models.Generate) error {
		return nil
	}

	generate := &models.Generate{
		Key:      "test_key",
		Table:    "test_table",
		Quantity: 50,
	}

	err := mockRepo.SaveDetailsGenerate(generate)

	assert.NoError(t, err)
}

func TestGenerateData(t *testing.T) {
	mockRepo := &mocks.MockGenerateRepository{}

	mockRepo.GenerateDataFunc = func(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error) {
		return []map[string]interface{}{
			{"name": "John Doe", "age": 30},
			{"name": "Jane Smith", "age": 25},
		}, nil
	}

	generateRequest := &request.GenerateRequest{
		Quantity: 100,
	}

	key := "test_key"

	results, err := mockRepo.GenerateData(key, generateRequest)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.True(t, len(results) <= 100)
}
