package repository

import (
	"github.com/respati123/money-tracking/internal/entity"
	"go.uber.org/zap"
)

type CategoryRepository struct {
	Repository[entity.Category]
	log *zap.Logger
}

func NewCategoryRepository(log *zap.Logger) *CategoryRepository {
	return &CategoryRepository{
		log: log,
	}
}
