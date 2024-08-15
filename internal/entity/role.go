package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID     string `gorm:"type:uuid;default:uuid_generate_v4()"`
	RoleCode int
	Name     string
	Alias    string
	Base
}
