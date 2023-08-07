package mocks

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/stretchr/testify/mock"
)

type MockGenerateService struct {
	mock.Mock
}

func (m *MockGenerateService) GetGenerateAll() ([]*models.Generate, error) {
	args := m.Called()
	return args.Get(0).([]*models.Generate), args.Error(1)
}

func (m *MockGenerateService) SaveDetailsGenerate(generate *models.Generate) error {
	args := m.Called(generate)
	return args.Error(0)
}

func (m *MockGenerateService) GenerateData(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error) {
	args := m.Called(key, generateRequest)
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}
