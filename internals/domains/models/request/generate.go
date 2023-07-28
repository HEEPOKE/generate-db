package request

import "github.com/HEEPOKE/generate-db/internals/domains/models"

type GenerateRequest struct {
	Table    string                            `json:"table"`
	Quantity int64                             `json:"quantity"`
	DataKey  string                            `json:"key"`
	Columns  map[string]models.GenerateOptions `json:"columns"`
}
