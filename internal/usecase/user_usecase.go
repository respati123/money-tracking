package usecase

import (
	"net/http"
	"strconv"

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

type UserUseCase struct {
	userRepo  *repository.UserRepository
	log       *zap.Logger
	db        *gorm.DB
	converter *converter.Converter
}

func (u *UserUseCase) CreateUser(ctx *gin.Context, request model.UserCreateRequest) model.ResponseInterface {
	tx := u.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err := u.userRepo.CountByEmail(tx, request.Email)
	if err != nil {
		u.log.Info("when count email")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusBadRequest,
			Message:    constants.EmailAlreadyExists,
		}
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		u.log.Info("hashing password")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusBadRequest,
			Message:    constants.EmailAlreadyExists,
		}
	}

	user := &entity.User{
		Email:       request.Email,
		Password:    hashedPassword,
		PhoneNumber: request.PhoneNumber,
		UserCode:    util.GenerateNumber(4),
	}

	err = u.userRepo.Create(tx, user)
	if err != nil {
		u.log.Info("when create user to database")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		u.log.Info("when transaction commit")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Message:    constants.Success,
	}
}

func (u *UserUseCase) DeleteUser(ctx *gin.Context, id string) model.ResponseInterface {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := u.userRepo.GetUserByID(tx, id)

	if err != nil {
		u.log.Info("when find user by id")

		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusBadRequest,
			Message:    constants.UserNotFound,
		}
	}

	err = u.userRepo.Delete(tx, user)

	if err != nil {
		u.log.Info("when deleting user")

		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		u.log.Info("when commit transaction")

		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Message:    constants.Success,
	}
}

func (u *UserUseCase) GetListUser(ctx *gin.Context, request model.PaginationRequest) model.ResponseInterface {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	users := new([]entity.User)

	_, metadata, err := u.userRepo.FindAll(tx, users, request)

	if err != nil {
		u.log.Info("error get list user")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	responses := model.PaginationResponse{
		Data:     u.converter.UserConverter.ToUserResponses(users),
		Metadata: metadata,
	}

	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Message:    constants.Success,
		Data:       responses,
	}

}

func (u *UserUseCase) GetUser(ctx *gin.Context, id int) model.ResponseInterface {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := u.userRepo.GetUserByID(tx, strconv.Itoa(id))
	if err != nil {
		u.log.Info("error get list user")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	response := u.converter.UserConverter.ToUserResponse(user)
	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Message:    constants.Success,
		Data:       response,
	}
}

func (u *UserUseCase) UpdateUser(ctx *gin.Context, id string, request model.UserUpdateRequest) model.ResponseInterface {
	tx := u.db.WithContext(ctx).Begin()
	tx.Rollback()

	user, err := u.userRepo.GetUserByID(tx, id)
	if err != nil {
		u.log.Info("when find user by uuid")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusBadRequest,
			Message:    constants.UserNotFound,
		}
	}

	user.Email = request.Email
	user.PhoneNumber = request.PhoneNumber

	err = u.userRepo.Update(tx, user)
	if err != nil {
		u.log.Info("when update user to database")
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Message:    constants.Success,
	}
}

func NewUserUsecase(
	db *gorm.DB,
	log *zap.Logger,
	converter *converter.Converter,
	userRepo *repository.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		userRepo:  userRepo,
		db:        db,
		converter: converter,
		log:       log,
	}
}
