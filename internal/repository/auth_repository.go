package repository

import (
	"fmt"

	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepository struct {
	log *zap.Logger
}

func (a *AuthRepository) Login(tx *gorm.DB, request model.LoginRequest) (*entity.User, error) {
	var user entity.User
	err := tx.Find(&user, "email = ?", request.Email).Error
	if err != nil {
		fmt.Print("pritn", err)
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepository) Register(tx *gorm.DB, user *entity.User) error {
	return tx.Create(&user).Error
}

func NewAuthRepository(log *zap.Logger) *AuthRepository {
	return &AuthRepository{
		log: log,
	}
}
