package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/model/converter"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionUsecase struct {
	db              *gorm.DB
	log             *zap.Logger
	converter       *converter.Converter
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(
	db *gorm.DB,
	log *zap.Logger,
	converter *converter.Converter,
	transactionRepo *repository.TransactionRepository,
) *TransactionUsecase {
	return &TransactionUsecase{
		db:              db,
		log:             log,
		converter:       converter,
		transactionRepo: *transactionRepo,
	}
}

func (tu *TransactionUsecase) Create(ctx *gin.Context, request model.TransactionRequest) model.ResponseInterface {
	user, _ := util.GetUserData(ctx)

	tx := tu.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var transaction = new(entity.Transaction)
	transaction.TransactionCode = uint(util.GenerateNumber(4))
	transaction.CategoryTypeCode = request.CategoryTypeCode
	transaction.Description = request.Description
	transaction.Title = request.Title
	transaction.Amount = request.Amount
	transaction.UserCode = uint(user.UserCode)
	transaction.TransactionType = request.TransactionType

	err := tu.transactionRepo.Create(tx, transaction)
	if err != nil {
		tu.log.Error("Error while create transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		tu.log.Error("Error while commit transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    "Success",
		Data:       "Successfully created the transaction",
		StatusCode: http.StatusOK,
	}
}

func (tu *TransactionUsecase) GetTransaction(ctx *gin.Context, code int) model.ResponseInterface {
	tx := tu.db.WithContext(ctx)

	var transaction = new(entity.Transaction)
	transactions, err := tu.transactionRepo.FindByCode(tx, transaction, "transaction_code", code)
	if err != nil {
		tu.log.Error("Error while get transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    "Success",
		Data:       tu.converter.ToTransactionResponse(transactions),
		StatusCode: http.StatusOK,
	}
}

func (tu *TransactionUsecase) UpdateTransaction(ctx *gin.Context, request model.TransactionUpdateRequest, id string) model.ResponseInterface {

	tx := tu.db.WithContext(ctx).Begin()

	var transaction = new(entity.Transaction)
	_, err := tu.transactionRepo.FindByField(tx, transaction, "uuid", id)

	if err != nil {
		tu.log.Error("Error while get transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	transaction.TransactionType = request.TransactionType
	transaction.CategoryTypeCode = request.CategoryTypeCode
	transaction.Description = request.Description
	transaction.Title = request.Title
	transaction.Amount = request.Amount

	err = tu.transactionRepo.Update(tx, transaction)
	if err != nil {
		tu.log.Error("Error while update transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		tu.log.Error("Error while commit transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    "Success",
		Data:       "Successfully updated the transaction",
		StatusCode: http.StatusOK,
	}

}

func (tu *TransactionUsecase) FindAll(ctx *gin.Context, pagination model.PaginationRequest) model.ResponseInterface {
	tx := tu.db.WithContext(ctx)

	var transactions = new([]entity.Transaction)
	transactions, metadata, err := tu.transactionRepo.FindAll(tx, transactions, pagination)

	if err != nil {
		tu.log.Error("Error while get transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "Error",
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	responses := model.PaginationResponse{
		Data:     tu.converter.ToTransactionResponses(transactions),
		Metadata: metadata,
	}

	return model.ResponseInterface{
		Message:    "Success",
		Data:       responses,
		StatusCode: http.StatusOK,
	}

}
