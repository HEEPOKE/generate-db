package examples

import "github.com/HEEPOKE/generate-db/pkg/enums"

type GenerateExample struct {
	Table    string                     `json:"table" validate:"required" example:"users"`
	Columns  map[string]GenerateOptions `json:"columns" validate:"required"`
	Quantity int64                      `json:"quantity" validate:"required" example:"1000"`
	DataKey  string                     `json:"key" example:""`
}

type GenerateOptions struct {
	Default      string         `json:"default" example:""`
	Types        enums.Category `json:"types"`
	Length       int            `json:"length"`
	AutoGenerate bool           `json:"gen"`
}
