package mocks

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/stretchr/testify/mock"
)

type MockGenerateRepository struct {
	mock.Mock
}

func (m *MockGenerateRepository) GetGenerateAll() ([]*models.Generate, error) {
	args := m.Called()
	return args.Get(0).([]*models.Generate), args.Error(1)
}

func (m *MockGenerateRepository) SaveDetailsGenerate(generate *models.Generate) error {
	args := m.Called(generate)
	return args.Error(0)
}

func (m *MockGenerateRepository) GenerateData(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error) {
	args := m.Called(key, generateRequest)
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}
