package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/util"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	db       *gorm.DB
	log      *logger.CustomLogger
	config   util.Config
	authRepo *repository.AuthRepository
	userRepo *repository.UserRepository
}

func (a *AuthUsecase) Login(ctx *gin.Context, request model.LoginRequest) (model.LoginResponse, error) {
	tx := a.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	user, err := a.authRepo.Login(tx, request)
	if err != nil {
		if err == constants.ErrUserNotFound {
			a.log.ErrorWithFields(ctx, "error user not found ", err)
			return model.LoginResponse{}, err
		}
		a.log.ErrorWithFields(ctx, "error user not found ", err)
		return model.LoginResponse{}, err
	}
	isValid := util.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		a.log.ErrorWithFields(ctx, "error check hashing password ", nil)
		return model.LoginResponse{}, err
	}
	jwtToken, expiredAt, err := util.GenerateJwtToken(util.JWTParams{
		Payload:    user.UUID,
		SecretKey:  a.config.SECRET_KEY_JWT,
		ExpireTime: a.config.JWT_EXPIRE_TIME,
	})
	if err != nil {
		a.log.ErrorWithFields(ctx, "error generate jwt token ", err)
		return model.LoginResponse{}, err
	}
	jwtRefreshToken, _, err := util.GenerateJwtToken(util.JWTParams{
		Payload:    user.UUID,
		SecretKey:  a.config.SECRET_KEY_JWT,
		ExpireTime: a.config.JWT_REFRESH_EXPIRE_TIME,
	})
	if err != nil {
		a.log.ErrorWithFields(ctx, "error generate jwt refresh token ", err)

		return model.LoginResponse{}, err
	}
	if err := tx.Commit().Error; err != nil {
		a.log.ErrorWithFields(ctx, "error commit transaction ", err)

		return model.LoginResponse{}, err
	}
	return model.LoginResponse{
		Token:        jwtToken,
		RefreshToken: jwtRefreshToken,
		ExpiredAt:    expiredAt.String(),
	}, nil
}

func (a *AuthUsecase) Register(ctx *gin.Context, request model.RegisterRequest) error {
	tx := a.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	user, err := a.userRepo.CountByEmail(tx, request.Email)
	if err == nil && user != 0 {
		a.log.ErrorWithFields(ctx, "email already exists ", err)

		return constants.ErrUserAlreadyExist
	}
	hashPassword, err := util.HashPassword(request.Password)
	if err != nil {
		a.log.ErrorWithFields(ctx, "error hashing password ", err)

		return err
	}
	var newUser entity.User
	newUser.Email = request.Email
	newUser.UserCode = util.GenerateNumber(4)
	newUser.Password = hashPassword
	newUser.PhoneNumber = request.PhoneNumber
	err = a.userRepo.Create(tx, &newUser)
	if err != nil {
		a.log.ErrorWithFields(ctx, "error create user ", err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		a.log.ErrorWithFields(ctx, "error commit transaction ", err)

		return err
	}
	return nil
}

func NewAuthUsecase(
	db *gorm.DB,
	log *logger.CustomLogger,
	config util.Config,
	authRepo *repository.AuthRepository,
	userRepo *repository.UserRepository,
) *AuthUsecase {
	return &AuthUsecase{
		db:       db,
		log:      log.Module("auth-service"),
		config:   config,
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
