package repository

import (
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	log *logger.CustomLogger
}

func NewUserRepository(log *logger.CustomLogger) *UserRepository {
	return &UserRepository{
		log: log.Module("user-repository"),
	}
}

func (a *UserRepository) FindByEmail(tx *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	err := tx.Take(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(db *gorm.DB, result *entity.User, id string) (*entity.User, error) {
	err := db.Where("uuid =?", id).First(&result).Error
	return result, err
}

func (r *UserRepository) FindByCode(db *gorm.DB, result *entity.User, id int) (*entity.User, error) {
	err := db.Where("user_code =?", id).First(&result).Error
	return result, err
}

func (r *UserRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Where("email = ?", email).Count(&total).Error
	return total, err
}
