package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
)

type TransactionController struct {
	log                *zap.Logger
	transactionUsecase *usecase.TransactionUsecase
}

func NewTransactionController(
	transactionUsecase *usecase.TransactionUsecase,
	log *zap.Logger,
) *TransactionController {
	return &TransactionController{
		log:                log,
		transactionUsecase: transactionUsecase,
	}
}

// Create Transaction
// @Summary Create Transaction
// @Description Create Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body model.TransactionRequest true "Transaction Body"
// @Success 200 {object} model.Response{response_data=string} "Successfully create transaction"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /transaction [post]
// @Security ApiKeyAuth
func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var request model.TransactionRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		tc.log.Error("Error while binding json", zap.Any("error", err.Error()))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while binding json", err)
		return
	}

	response := tc.transactionUsecase.Create(ctx, request)
	util.Response(ctx, response)
}

// Get Transaction
// @Summary Get Transaction
// @Description Get Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction_code path int true "Transaction Code"
// @Success 200 {object} model.Response{response_data=model.TransactionResponse} "Successfully get transaction"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /transaction/{transaction_code} [get]
// @Security ApiKeyAuth
func (tc *TransactionController) GetTransaction(ctx *gin.Context) {
	paramId := ctx.Param("transaction_code")
	code, _ := strconv.Atoi(paramId)
	response := tc.transactionUsecase.GetTransaction(ctx, code)
	util.Response(ctx, response)
}

// Update Transaction
// @Summary Update Transaction
// @Description Update Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param uuid path string true "uuid"
// @Param transaction body model.TransactionUpdateRequest true "Transaction Body"
// @Success 200 {object} model.Response{response_data=string} "Successfully update transaction"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /transaction/{uuid} [put]
// @Security ApiKeyAuth
func (tc *TransactionController) UpdateTransaction(ctx *gin.Context) {
	var requestUpdate model.TransactionUpdateRequest
	paramId := ctx.Param("uuid")

	if err := ctx.ShouldBindJSON(&requestUpdate); err != nil {
		tc.log.Error("Error while binding json", zap.Any("error", err.Error()))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while binding json", err)
		return
	}

	response := tc.transactionUsecase.UpdateTransaction(ctx, requestUpdate, paramId)
	util.Response(ctx, response)
}

// Find All Transaction
// @Summary Find All Transaction
// @Description Find All Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param pagination body model.PaginationRequest true "Pagination Request"
// @Success 200 {object} model.Response{response_data=model.PaginationResponse{data=[]model.TransactionResponse}} "Successfully find all transaction"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /transaction/all [post]
// @Security ApiKeyAuth
func (tc *TransactionController) FindAll(ctx *gin.Context) {
	var request model.PaginationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		tc.log.Error("Error while binding json", zap.Any("error", err.Error()))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while binding json", err)
		return
	}
	response := tc.transactionUsecase.FindAll(ctx, request)
	util.Response(ctx, response)
}
