package converter

import (
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
)

type TransactionConverter struct{}

func NewTransactionConverter() *TransactionConverter {
	return &TransactionConverter{}
}

func (tc *TransactionConverter) ToTransactionResponse(transaction *entity.Transaction) *model.TransactionResponse {
	return &model.TransactionResponse{
		ID:               transaction.ID,
		UUID:             transaction.UUID,
		TransactionCode:  transaction.TransactionCode,
		CategoryTypeCode: transaction.CategoryTypeCode,
		Description:      transaction.Description,
		Title:            transaction.Title,
		Amount:           transaction.Amount,
		CreatedAt:        transaction.CreatedAt.String(),
		UpdatedAt:        transaction.UpdatedAt.String(),
		DeletedAt:        transaction.DeletedAt.Time.String(),
		CreatedBy:        transaction.CreatedBy,
		UpdatedBy:        transaction.UpdatedBy,
		DeletedBy:        transaction.DeletedBy,
	}

}

func (tc *TransactionConverter) ToTransactionResponses(transactions *[]entity.Transaction) *[]model.TransactionResponse {
	var transactionResponses []model.TransactionResponse
	for _, transaction := range *transactions {
		transactionResponses = append(transactionResponses, *tc.ToTransactionResponse(&transaction))
	}
	return &transactionResponses
}

func (tc *TransactionConverter) ToTransactionType(transactionType string) constants.TransactionType {
	switch transactionType {
	case "debit":
		return constants.TransactionTypes.Debit
	case "credit":
		return constants.TransactionTypes.Credit
	default:
		return constants.TransactionTypes.Debit
	}
}
