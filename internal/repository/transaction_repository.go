package repository

import (
	"github.com/respati123/money-tracking/internal/entity"
	"go.uber.org/zap"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
	log *zap.Logger
}

func NewTransactionRepository(log *zap.Logger) *TransactionRepository {
	return &TransactionRepository{
		log: log,
	}
}
