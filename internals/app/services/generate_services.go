package services

import (
	"github.com/HEEPOKE/generate-db/internals/core/interfaces"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
)

type GenerateService struct {
	generateRepository interfaces.GenerateRepository
}

func NewGenerateService(generateRepository interfaces.GenerateRepository) *GenerateService {
	return &GenerateService{generateRepository: generateRepository}
}

func (s *GenerateService) GetGenerateAll() ([]*models.Generate, error) {
	return s.generateRepository.GetGenerateAll()
}

func (s *GenerateService) SaveDetailsGenerate(generate *models.Generate) error {
	return s.generateRepository.SaveDetailsGenerate(generate)
}

func (s *GenerateService) GenerateData(data models.Generate, generateRequest *request.GenerateRequest) (models.JsonStructure, error) {
	return s.generateRepository.GenerateData(data, generateRequest)
}
