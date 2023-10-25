package interfaces

import "github.com/HEEPOKE/generate-db/internals/domains/models"

type UtilitiesRepository interface {
	CheckKeyData(key string) (models.JsonStructure, error)
}
