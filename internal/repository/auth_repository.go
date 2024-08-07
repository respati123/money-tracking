package repository

import (
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	log *logger.CustomLogger
}

func (a *AuthRepository) Login(tx *gorm.DB, request model.LoginRequest) (*entity.User, error) {
	var user entity.User
	err := tx.Take(&user, "email = ?", request.Email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepository) Register(tx *gorm.DB, user *entity.User) error {
	err := tx.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func NewAuthRepository(log *logger.CustomLogger) *AuthRepository {
	return &AuthRepository{
		log: log,
	}
}
