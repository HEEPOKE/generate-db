package examples

import (
	"time"

	"github.com/HEEPOKE/generate-db/pkg/enums"
)

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

type GenerateResponseExample struct {
	ID          uint      `json:"id" example:"1"`
	Key         string    `json:"key" example:"123456"`
	Table       string    `json:"table" example:"users"`
	Quantity    int64     `json:"quantity" example:"1000"`
	TimeExpired time.Time `json:"time_expired" example:"2022-01-01 00:00:00"`
	CreatedAt   time.Time `json:"created_at" example:"2022-01-01 00:00:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2022-01-01 00:00:00"`
	DeletedAt   time.Time `json:"deleted_at" example:"0001-01-01 00:00:00"`
}
