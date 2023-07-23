package interfaces

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
)

type GenerateRepository interface {
	GetGenerateAll() ([]*models.Generate, error)
	SaveDetailsGenerate(generate *models.Generate) error
	GenerateData(generateRequest *request.GenerateRequest) (interface{}, error)
}
