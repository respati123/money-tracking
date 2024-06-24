package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Age         string
	Description string
}
