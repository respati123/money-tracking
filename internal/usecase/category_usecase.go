package usecase

import (
	"net/http"
	"strings"

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

type CategoryUseCase struct {
	db           *gorm.DB
	log          *zap.Logger
	categoryRepo repository.CategoryRepository
	converter    *converter.Converter
}

func NewCategoryUsecase(db *gorm.DB, log *zap.Logger, converter *converter.Converter, categoryRepo *repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		db:           db,
		log:          log,
		categoryRepo: *categoryRepo,
		converter:    converter,
	}
}

func (cu *CategoryUseCase) Create(ctx *gin.Context, request model.CategoryCreateRequest) model.ResponseInterface {
	tx := cu.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var category = new(entity.Category)
	category.CategoryCode = util.GenerateNumber(4)
	category.Alias = request.Alias
	category.Name = request.Name

	err := cu.categoryRepo.Create(tx, category)

	if err != nil {
		if strings.Contains(err.Error(), constants.DuplicateKey) {
			cu.log.Error("Error while create category", zap.Any("error", err.Error()))
			return model.ResponseInterface{
				Message:    constants.Error,
				Error:      constants.ErrDuplicate("category"),
				StatusCode: http.StatusBadRequest,
			}
		}
		cu.log.Error("Error while create category", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    constants.Error,
			Error:      constants.ErrInternalServerError,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		cu.log.Error("Error while commit transaction", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    constants.Error,
			Error:      gorm.ErrInvalidTransaction,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		Data:       "Successfully created the category",
		StatusCode: http.StatusOK,
	}
}

func (cu *CategoryUseCase) Delete(ctx *gin.Context, id string) model.ResponseInterface {
	tx := cu.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var category = new(entity.Category)
	_, err := cu.categoryRepo.FindByField(tx, category, "uuid", id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cu.log.Error("Error while find by field", zap.Error(err))
			return model.ResponseInterface{
				Message:    constants.Error,
				StatusCode: http.StatusBadRequest,
				Error:      constants.ErrNotFound("category uuid"),
			}
		}
		cu.log.Error("Error while find by field", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
		}
	}

	err = cu.categoryRepo.Delete(tx, category)

	if err != nil {
		cu.log.Error("Error while delete category", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		cu.log.Error("Error while commit transaction", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		StatusCode: http.StatusOK,
		Data:       "Successfully deleted the category",
	}

}

func (cu *CategoryUseCase) Update(ctx *gin.Context, request model.CategoryUpdateRequest, id string) model.ResponseInterface {
	tx := cu.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var category = new(entity.Category)
	_, err := cu.categoryRepo.FindByField(tx, category, "uuid", id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cu.log.Error("Error while find by field uuid", zap.Error(err))
			return model.ResponseInterface{
				Message:    constants.Error,
				StatusCode: http.StatusBadRequest,
				Error:      constants.ErrNotFound("category "),
			}
		}
		cu.log.Error("Error while commit transaction", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
		}
	}

	category.Alias = request.Alias
	category.Name = request.Name

	err = cu.categoryRepo.Update(tx, category)

	if err != nil {
		cu.log.Error("Error while update category", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		cu.log.Error("Error while commit transaction", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		StatusCode: http.StatusOK,
		Data:       "Successfully updated the category",
	}
}

func (cu *CategoryUseCase) FindAll(ctx *gin.Context, pagination model.PaginationRequest) model.ResponseInterface {
	tx := cu.db.WithContext(ctx)
	var categories []entity.Category

	_, metadata, err := cu.categoryRepo.FindAll(tx, &categories, pagination)

	if err != nil {
		cu.log.Error("Error while find all category", zap.Error(err))
		return model.ResponseInterface{
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
			Message:    constants.Error,
		}
	}

	response := model.PaginationResponse{
		Data:     cu.converter.ToCategoryResponses(&categories),
		Metadata: metadata,
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		StatusCode: http.StatusOK,
		Data:       response,
	}
}

func (cu *CategoryUseCase) GetCategory(ctx *gin.Context, id int) model.ResponseInterface {
	tx := cu.db.WithContext(ctx)
	var category entity.Category

	_, err := cu.categoryRepo.FindByCode(tx, &category, "category_code", id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cu.log.Error("Error while not found", zap.Error(err))
			return model.ResponseInterface{
				StatusCode: http.StatusBadRequest,
				Error:      constants.ErrNotFound("category"),
				Message:    constants.Error,
			}
		}
		cu.log.Error("Error while find category by code", zap.Error(err))
		return model.ResponseInterface{
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrInternalServerError,
			Message:    constants.Error,
		}
	}

	return model.ResponseInterface{
		Data:       cu.converter.ToCategoryResponse(category),
		StatusCode: http.StatusOK,
		Message:    constants.Success,
	}
}
