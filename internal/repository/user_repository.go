package repository

import (
	"context"

	"github.com/respati123/money-tracking/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserIDByEmail(ctx context.Context, email string) (uint, error)
}

type userRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func (u *userRepository) GetUserIDByEmail(ctx context.Context, email string) (uint, error) {
	var user entity.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	return u.db.Delete(&entity.User{}, id).Error
}

func (u *userRepository) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}
