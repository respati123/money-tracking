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

type RoleUsecase struct {
	roleRepository *repository.RoleRepository
	log            *zap.Logger
	db             *gorm.DB
	response       model.ResponseInterface
	converter      *converter.Converter
}

func NewRoleUsecase(db *gorm.DB, log *zap.Logger, converter *converter.Converter, roleRepository *repository.RoleRepository) *RoleUsecase {
	return &RoleUsecase{
		roleRepository: roleRepository,
		log:            log,
		db:             db,
		response:       model.ResponseInterface{},
		converter:      converter,
	}
}

func (rs *RoleUsecase) Create(ctx *gin.Context, request model.RoleCreateRequest) model.ResponseInterface {
	tx := rs.db.WithContext(ctx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var role = new(entity.Role)
	_, err := rs.roleRepository.FindByField(tx, role, "name", request.Name)

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

	tx = tx.Begin()

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
		Message:    constants.Success,
		Data:       "Successfully created the role",
		StatusCode: http.StatusOK,
	}
}

func (rs *RoleUsecase) Update(ctx *gin.Context, request model.RoleUpdateRequest, uuid string) model.ResponseInterface {
	tx := rs.db.WithContext(ctx)

	var role = new(entity.Role)
	_, err := rs.roleRepository.FindByField(tx, role, "uuid", uuid)
	if err != nil {
		rs.log.Error("error find count by id", zap.Any("error", err.Error()))
		return model.ResponseInterface{
			Message:    "error role services",
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	if role.RoleCode == 0 {
		rs.log.Error("error find count by id", zap.Any("error", constants.ErrNotFound("role")))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      constants.ErrNotFound("role"),
		}
	}

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
		if strings.Contains(err.Error(), constants.DuplicateKey) {
			tx.Rollback()
			rs.log.Error("error while update role", zap.Any("error", err.Error()))
			return model.ResponseInterface{
				Message:    constants.Error,
				StatusCode: http.StatusBadRequest,
				Error:      constants.ErrDuplicate("role name"),
			}
		}

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
	tx := rs.db.WithContext(ctx).Begin()

	var role = new(entity.Role)
	_, err := rs.roleRepository.FindByField(tx, role, "uuid", id)
	if err != nil {
		rs.log.Error("Error while count by id", zap.Any("error", err.Error()))

		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	err = rs.roleRepository.Delete(tx, role)

	if err != nil {
		tx.Rollback()
		rs.log.Error("error while delete role", zap.Any("error", err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	if err := tx.Commit().Error; err != nil {
		rs.log.Error("error while transaction commit", zap.Any("error", err))
		return model.ResponseInterface{
			Message:    "error while transaction commit",
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		StatusCode: http.StatusOK,
		Data:       "Successfully deleted the role",
	}
}

func (rs *RoleUsecase) GetRole(ctx *gin.Context, id int) model.ResponseInterface {
	tx := rs.db.WithContext(ctx)

	var role = new(entity.Role)
	_, err := rs.roleRepository.FindByCode(tx, role, "role_code", id)

	if err != nil {
		rs.log.Error("Error while find role by code", zap.Any("error", err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusBadRequest,
			Error:      constants.ErrNotFound("role"),
		}
	}

	return model.ResponseInterface{
		Message:    constants.Success,
		StatusCode: http.StatusOK,
		Data:       role,
	}
}

func (rs *RoleUsecase) FindAll(ctx *gin.Context, pagination model.PaginationRequest) model.ResponseInterface {
	tx := rs.db.WithContext(ctx)

	var roles = new([]entity.Role)
	_, metadata, err := rs.roleRepository.FindAll(tx, roles, pagination)
	if err != nil {
		rs.log.Error("Error while get find all role", zap.Any("error", err))
		return model.ResponseInterface{
			Message:    constants.Error,
			StatusCode: http.StatusOK,
			Error:      err,
		}
	}

	response := model.PaginationResponse{
		Data:     rs.converter.RoleConverter.ToRoleResponses(roles),
		Metadata: metadata,
	}

	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Data:       response,
		Message:    constants.Success,
	}
}
