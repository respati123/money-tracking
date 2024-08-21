package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UUID             string `gorm:"type:uuid;default:uuid_generate_v4()"`
	TransactionCode  uint
	CategoryTypeCode uint
	UserCode         uint
	// Category         Category `gorm:"foreignKey:category_type_code;references:category_type_code"`
	TransactionType string
	Description     string
	Title           string
	Amount          float64
	Base
}

func (Transaction) TableName() string {
	return "transaction"
}
