package interfaces

import "github.com/HEEPOKE/generate-db/internals/domains/models"

type GenerateRepository interface {
	GetGenerateAll() ([]*models.Generate, error)
	SaveDetailsGenerate(generate *models.Generate) error
}
