package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UUID             string `gorm:"type:uuid;default:uuid_generator_v4()"`
	CategoryTypeCode int
	Name             string
	Alias            string
	Base
}

func (Category) TableName() string {
	return "category_type"
}
