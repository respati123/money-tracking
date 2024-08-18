package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID
	UserCode    int
	Email       string
	Password    string
	PhoneNumber string
	Base
}
