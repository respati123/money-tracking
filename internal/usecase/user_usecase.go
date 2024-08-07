package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/model/converter"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/util"
	"gorm.io/gorm"
)

type UserUseCase struct {
	userRepo  *repository.UserRepository
	converter *converter.Converter
	log       *logger.CustomLogger
	db        *gorm.DB
}

func (u *UserUseCase) CreateUser(ctx *gin.Context, request model.UserCreateRequest) error {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	_, err := u.userRepo.CountByEmail(tx, request.Email)
	if err != nil {
		u.log.ErrorWithFields(ctx, "email already exists", err)
		return err
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		u.log.ErrorWithFields(ctx, "error hashing password", err)
		return err
	}

	user := &entity.User{
		Email:       request.Email,
		Password:    hashedPassword,
		PhoneNumber: request.PhoneNumber,
		UserCode:    util.GenerateNumber(4),
	}

	err = u.userRepo.Create(tx, user)
	if err != nil {
		u.log.ErrorWithFields(ctx, "error create user", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		u.log.ErrorWithFields(ctx, "error transaction commit", err)
		return err
	}

	return nil
}

func (u *UserUseCase) DeleteUser(ctx *gin.Context, id string) error {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	_, err := u.userRepo.FindByID(tx, user, id)

	if err != nil {
		u.log.ErrorWithFields(ctx, "user not found", err)

		return err
	}

	err = u.userRepo.Delete(tx, user)

	if err != nil {
		u.log.ErrorWithFields(ctx, "error deleting user", err)

		return err
	}

	if err := tx.Commit().Error; err != nil {
		u.log.ErrorWithFields(ctx, "error commit transaction", err)

		return err
	}

	return nil
}

func (u *UserUseCase) GetListUser(ctx *gin.Context, request model.PaginationRequest) (*model.PaginationResponse, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	users := new([]entity.User)

	_, metadata, err := u.userRepo.FindAll(tx, users, request)

	if err != nil {
		u.log.ErrorWithFields(ctx, "error get list user", err)

		return nil, err
	}

	responses := model.PaginationResponse{
		Data:     u.converter.UserConverter.ToResponseUsers(users),
		Metadata: metadata,
	}

	return &responses, nil

}

func (u *UserUseCase) GetUser(ctx *gin.Context, id int) (*model.UserResponse, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	_, err := u.userRepo.FindByID(tx, user, strconv.Itoa(id))
	if err != nil {
		u.log.ErrorWithFields(ctx, "error get list user", err)

		return nil, err
	}

	response := u.converter.UserConverter.ToResponseUser(user)
	return response, nil
}

func (u *UserUseCase) UpdateUser(ctx *gin.Context, id string, request model.UserUpdateRequest) error {
	tx := u.db.WithContext(ctx).Begin()
	tx.Rollback()

	var user = new(entity.User)
	_, err := u.userRepo.FindByID(tx, user, id)
	if err != nil {
		u.log.ErrorWithFields(ctx, "error user not found", err)

		return err
	}

	user.Email = request.Email
	user.PhoneNumber = request.PhoneNumber

	// err = u.userRepo.Update(tx, user)
	// if err != nil {
	// u.log.WithFields(logrus.Fields{
	// 	"trace_id": ct
	// })
	// }

	return nil
}

func NewUserUsecase(
	db *gorm.DB,
	log *logger.CustomLogger,
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
