package repository

import (
	"context"

	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(ctx context.Context, request model.LoginRequest) (*entity.User, error)
	Register(ctx context.Context, user *entity.User) error
}

type authRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func (a *authRepository) Login(ctx context.Context, request model.LoginRequest) (*entity.User, error) {
	var user entity.User
	err := a.db.WithContext(ctx).Where("email = ?", request.Email).First(&user)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return nil, constants.ErrUserNotFound
		}
		return nil, err.Error
	}
	return &user, nil
}

func (a *authRepository) Register(ctx context.Context, user *entity.User) error {
	return a.db.WithContext(ctx).Create(&user).Error
}

func NewAuthRepository(db *gorm.DB, log *logrus.Logger) AuthRepository {
	return &authRepository{
		db:  db,
		log: log,
	}
}
