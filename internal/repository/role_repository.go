package repository

import (
	"github.com/respati123/money-tracking/internal/entity"
	"go.uber.org/zap"
)

type RoleRepository struct {
	Repository[entity.Role]
	log *zap.Logger
}

func NewRoleRepository(log *zap.Logger) *RoleRepository {
	return &RoleRepository{
		log: log,
	}
}
