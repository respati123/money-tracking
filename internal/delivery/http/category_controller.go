package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
)

type CategoryController struct {
	log             *zap.Logger
	categoryUsecase usecase.CategoryUseCase
}

func NewCategoryController(categoryUsecase *usecase.CategoryUseCase, log *zap.Logger) *CategoryController {
	return &CategoryController{
		log:             log,
		categoryUsecase: *categoryUsecase,
	}
}

// Create Category
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Produce json
// @Accept json
// @Param body body model.CategoryCreateRequest true "create category request"
// @Success 200 {object} model.Response{response_data=string} "successfully find all role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /category/ [post]
// @Security ApiKeyAuth
func (cc *CategoryController) Create(ctx *gin.Context) {

	var request model.CategoryCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		cc.log.Error("Error while binding json", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, constants.Error, err)
		return
	}

	response := cc.categoryUsecase.Create(ctx, request)
	util.Response(ctx, response)
}

// Delete Category
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Produce json
// @Param uuid path string true "category uuid"
// @Success 200 {object} model.Response{response_data=string} "successfully delete role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /category/{uuid} [delete]
// @Security ApiKeyAuth
func (cc *CategoryController) Delete(ctx *gin.Context) {

	paramId := ctx.Param("uuid")

	response := cc.categoryUsecase.Delete(ctx, paramId)
	util.Response(ctx, response)
}

// FindAll Category
// @Summary Find All Category
// @Description Find All Category
// @Tags Category
// @Produce json
// @Param body body model.PaginationRequest true "pagination request"
// @Success 200 {object} model.Response{response_data=model.PaginationResponse{data=[]model.CategoryResponse}} "successfully find all role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /category/all [post]
// @Security ApiKeyAuth
func (cc *CategoryController) FindAll(ctx *gin.Context) {
	var request model.PaginationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		cc.log.Error("Error while binding json", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, constants.Error, err)
		return
	}

	response := cc.categoryUsecase.FindAll(ctx, request)
	util.Response(ctx, response)
}

// Update Category
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Produce json
// @Accept json
// @Param uuid path string true "category uuid"
// @Param body body model.CategoryUpdateRequest true "update category request"
// @Success 200 {object} model.Response{response_data=string} "successfully update role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /category/{uuid} [put]
// @Security ApiKeyAuth
func (cc *CategoryController) Update(ctx *gin.Context) {
	var paramId string
	var request model.CategoryUpdateRequest

	paramId = ctx.Param("uuid")
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		cc.log.Error("Error while binding json", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, constants.Error, err)
		return
	}

	response := cc.categoryUsecase.Update(ctx, request, paramId)
	util.Response(ctx, response)
}

// Find Category
// @Summary Find Category
// @Description Find Category
// @Tags Category
// @Produce json
// @Param category_code path int true "category code"
// @Success 200 {object} model.Response{response_data=model.CategoryResponse} "successfully find role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /category/{category_code} [get]
// @Security ApiKeyAuth
func (cc *CategoryController) Find(ctx *gin.Context) {

	paramId := ctx.Param("category_code")
	categoryCode, err := strconv.Atoi(paramId)
	if err != nil {
		cc.log.Error("Error while convert string to int", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, constants.Error, err)
		return
	}
	response := cc.categoryUsecase.GetCategory(ctx, categoryCode)
	util.Response(ctx, response)
}
