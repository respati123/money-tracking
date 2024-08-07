package entity

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Base struct {
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
