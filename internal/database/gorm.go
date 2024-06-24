package database

import (
	"github.com/respati123/money-tracking/internal/entity"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB, dbTest *gorm.DB) error {
	var migrations = []interface{}{entity.User{}}

	err := db.AutoMigrate(
		migrations...,
	)

	if err != nil {
		return err
	}
	err = dbTest.AutoMigrate(
		migrations...,
	)

	if err != nil {
		return err
	}

	return nil
}
