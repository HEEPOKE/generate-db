package services_test

import (
	"testing"
	"time"

	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/HEEPOKE/generate-db/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetGenerateAll(t *testing.T) {
	mockRepo := new(mocks.MockGenerateRepository)

	service := services.NewGenerateService(mockRepo)

	expectedData := []*models.Generate{
		{ID: 1, Key: "Generate 1", Table: "Content 1", Quantity: 100, TimeExpired: time.Now()},
		{ID: 2, Key: "Generate 2", Table: "Content 2", Quantity: 200, TimeExpired: time.Now()},
	}

	mockRepo.On("GetGenerateAll").Return(expectedData, nil)

	data, err := service.GetGenerateAll()

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, expectedData, data)

	mockRepo.AssertExpectations(t)
}

func TestSaveDetailsGenerate(t *testing.T) {
	mockRepo := new(mocks.MockGenerateRepository)

	service := services.NewGenerateService(mockRepo)

	generate := &models.Generate{
		ID:          1,
		Key:         "Generate 1",
		Table:       "Content 1",
		Quantity:    100,
		TimeExpired: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}

	mockRepo.On("SaveDetailsGenerate", generate).Return(nil)

	err := service.SaveDetailsGenerate(generate)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGenerateData(t *testing.T) {
	mockRepo := new(mocks.MockGenerateRepository)

	service := services.NewGenerateService(mockRepo)

	key := "some_key"
	generateRequest := &request.GenerateRequest{}

	expectedData := []map[string]interface{}{
		{"key": "value1"},
		{"key": "value2"},
	}

	mockRepo.On("GenerateData", key, generateRequest).Return(expectedData, nil)

	data, err := service.GenerateData(key, generateRequest)

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, expectedData, data)

	mockRepo.AssertExpectations(t)
}
