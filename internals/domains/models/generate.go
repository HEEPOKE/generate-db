package models

import (
	"time"

	"github.com/HEEPOKE/generate-db/pkg/enums"
	"gorm.io/gorm"
)

type Generate struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Key         string         `gorm:"type:VARCHAR(255);unique" json:"key"`
	Table       string         `gorm:"type:VARCHAR(255);" json:"table"`
	Quantity    int64          `gorm:"type:BIGINT" json:"quantity"`
	TimeExpired time.Time      `gorm:"type:timestamp" json:"time_expired"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type GenerateOptions struct {
	Default      string         `json:"default"`
	Types        enums.Category `json:"types"`
	Length       int            `json:"length"`
	AutoGenerate bool           `json:"gen"`
}
