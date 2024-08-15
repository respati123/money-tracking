package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleUsecase struct {
	roleRepository *repository.RoleRepository
	log            *zap.Logger
	db             *gorm.DB
	response       model.ResponseInterface
}

func NewRoleUsecase(db *gorm.DB, log *zap.Logger, roleRepository *repository.RoleRepository) *RoleUsecase {
	return &RoleUsecase{
		roleRepository: roleRepository,
		log:            log,
		db:             db,
		response:       model.ResponseInterface{},
	}
}

func (rs *RoleUsecase) Create(ctx *gin.Context, request model.RoleCreateRequest) model.ResponseInterface {
	tx := rs.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var role = new(entity.Role)
	err := rs.roleRepository.FindByField(tx, role, "name", request.Name)

	if err != nil && err != gorm.ErrRecordNotFound {
		rs.log.Info("name already exists", zap.Error(err))
		return model.ResponseInterface{
			Message:    "name already exists",
			Error:      err,
			StatusCode: http.StatusBadRequest,
		}
	}

	role.RoleCode = util.GenerateNumber(4)
	role.Name = request.Name
	role.Alias = request.Alias

	err = rs.roleRepository.Create(tx, role)
	if err != nil {
		if err != gorm.ErrDuplicatedKey {
			tx.Rollback()
			rs.log.Info("when create role", zap.Error(err))
			return model.ResponseInterface{
				Message:    "error duplicate name",
				Error:      err,
				StatusCode: http.StatusInternalServerError,
			}
		}
		tx.Rollback()
		rs.log.Info("when create role", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.InternalServerError,
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		rs.log.Info("when create role", zap.Error(err))
		return model.ResponseInterface{
			Message:    constants.InternalServerError,
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return model.ResponseInterface{
		Message:    "role created",
		StatusCode: http.StatusOK,
	}
}

func (rs *RoleUsecase) Update(ctx *gin.Context, request model.RoleUpdateRequest, uuid string) model.ResponseInterface {
	tx := rs.db.WithContext(ctx)

	count, err := rs.roleRepository.CountById(tx, uuid)
	if err != nil {
		rs.log.Error("error find count by id", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "error role services",
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	if count == 0 {
		rs.log.Error("role id not found", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "role id not found",
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	var role = new(entity.Role)
	role.Name = request.Name
	role.Alias = request.Alias
	role.RoleCode = util.GenerateNumber(4)

	tx = tx.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = rs.roleRepository.Update(tx, role)

	if err != nil {
		tx.Rollback()
		rs.log.Error("error while update role", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "error update role",
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	if err := tx.Commit().Error; err != nil {
		rs.log.Error("error while transaction commit", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "error while transaction commit",
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		StatusCode: http.StatusOK,
		Data:       "successfully updated the role.",
	}

}

func (rs *RoleUsecase) Delete(ctx *gin.Context, id string) model.ResponseInterface {
	tx := rs.db.WithContext(ctx)

	var role = new(entity.Role)
	err := rs.roleRepository.FindByField(tx, role, "uuid", id)
	if err != nil {
		rs.log.Error("Error while count by id", zap.Any("error", err.Error()))
		rs.response.Message = constants.Error
		rs.response.Error = err
		rs.response.StatusCode = http.StatusInternalServerError
		return rs.response
	}

	tx.Begin()

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	err = rs.roleRepository.Delete(tx, role)

	if err != nil {
		rs.log.Error("error while delete role", zap.Any("error", err.Error()))
		rs.response.Message = constants.Error
		rs.response.StatusCode = http.StatusInternalServerError
		rs.response.Error = err
		return rs.response
	}

	if err := tx.Commit().Error; err != nil {
		rs.log.Error("error while transaction commit", zap.Any("error", err.Error()))
		rs.response.Message = "error while transaction commit"
		rs.response.StatusCode = http.StatusInternalServerError
		rs.response.Error = err

		return rs.response
	}

	rs.response.Message = constants.Success
	rs.response.StatusCode = http.StatusOK
	rs.response.Data = "Successfully deleted the role"

	return rs.response
}
