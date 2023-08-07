package mocks

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
)

type MockGenerateRepository struct {
	GetGenerateAllFunc      func() ([]*models.Generate, error)
	SaveDetailsGenerateFunc func(generate *models.Generate) error
	GenerateDataFunc        func(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error)
}

func (m *MockGenerateRepository) GetGenerateAll() ([]*models.Generate, error) {
	return m.GetGenerateAllFunc()
}

func (m *MockGenerateRepository) SaveDetailsGenerate(generate *models.Generate) error {
	return m.SaveDetailsGenerateFunc(generate)
}

func (m *MockGenerateRepository) GenerateData(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error) {
	return m.GenerateDataFunc(key, generateRequest)
}
