package examples

import "github.com/HEEPOKE/generate-db/internals/domains/models"

type GenerateExample struct {
	Table    string                            `json:"table" validate:"required" example:"users"`
	Columns  map[string]models.GenerateOptions `json:"columns" validate:"required"`
	Quantity int64                             `json:"quantity" validate:"required" example:"1000"`
}
