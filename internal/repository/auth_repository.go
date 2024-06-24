package repository

import (
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(payload interface{})
}

type AuthRepositoryImpl struct {
	db gorm.DB
}

// Login implements AuthRepository.
func (repository *AuthRepositoryImpl) Login(payload interface{}) {
	panic("unimplemented")
}

func NewAuthRepository(db gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}
