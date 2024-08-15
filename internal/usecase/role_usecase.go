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
}

func NewRoleUsecase(db *gorm.DB, log *zap.Logger, roleRepository *repository.RoleRepository) *RoleUsecase {
	return &RoleUsecase{
		roleRepository: roleRepository,
		log:            log,
		db:             db,
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

	// rs.log.Info("role created", zap.Any("role", role))
	return model.ResponseInterface{
		Message:    "role created",
		StatusCode: http.StatusOK,
	}

}
