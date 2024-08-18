package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UUID         string `gorm:"type:uuid;default:uuid_generator_v4()"`
	CategoryCode int    `gorm:"unique;not null;column:transaction_type_code"`
	Name         string
	Alias        string
	Base
}

func (Category) TableName() string {
	return "category_type"
}
