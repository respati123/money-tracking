package usecase

import (
	"context"

	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase interface {
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userUseCase struct {
	userRepo repository.UserRepository
	log      *logrus.Logger
	db       *gorm.DB
}

func (u *userUseCase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	panic("unimplemented")
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetUserByID implements UserService.
func (u *userUseCase) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userUseCase) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserService(userRepo repository.UserRepository, db *gorm.DB, log *logrus.Logger) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
		db:       db,
		log:      log,
	}
}
