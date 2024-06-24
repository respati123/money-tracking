package service

import "github.com/respati123/money-tracking/internal/repository"

type AuthServiceImpl struct {
	repository repository.AuthRepository
}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(payload interface{}) {
	panic("unimplemented")
}

type AuthService interface {
	Login(payload interface{})
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		repository: repository,
	}
}
