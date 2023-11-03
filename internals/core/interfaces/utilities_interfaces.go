package interfaces

import "github.com/HEEPOKE/generate-db/internals/domains/models"

type UtilitiesRepository interface {
	CheckKeyData(key string) (models.JsonStructure, error)
	UpdatePathFileJson(table, key, path string) error
}
