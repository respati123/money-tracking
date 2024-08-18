package repository

import (
	"github.com/respati123/money-tracking/internal/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	log *zap.Logger
}

func (u *UserRepository) GetUserByID(tx *gorm.DB, id string) (*entity.User, error) {
	var user entity.User
	err := tx.Where("uuid =?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	err := tx.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) CountByEmail(tx *gorm.DB, email string) (int64, error) {
	var total int64
	err := tx.Model(&entity.User{}).Where("email = ?", email).Count(&total).Error
	return total, err
}

func NewUserRepository(log *zap.Logger) *UserRepository {
	return &UserRepository{
		log: log,
	}
}
