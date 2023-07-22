package models

import (
	"time"

	"gorm.io/gorm"
)

type Generate struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Key       string         `gorm:"type:VARCHAR(255);unique" json:"key"`
	Expired   int            `gorm:"type:VARCHAR(255)" json:"expired"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
