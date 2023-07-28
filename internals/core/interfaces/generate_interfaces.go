package interfaces

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
)

type GenerateRepository interface {
	GetGenerateAll() ([]*models.Generate, error)
	SaveDetailsGenerate(generate *models.Generate) error
	GenerateData(key string, generateRequest *request.GenerateRequest) ([]map[string]interface{}, error)
}
