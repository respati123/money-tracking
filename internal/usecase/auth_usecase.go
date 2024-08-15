package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	db       *gorm.DB
	log      *zap.Logger
	config   util.Config
	redis    *redis.Client
	authRepo *repository.AuthRepository
	userRepo *repository.UserRepository
}

func (a *AuthUsecase) Login(ctx *gin.Context, request model.LoginRequest) model.ResponseInterface {
	tx := a.db.WithContext(ctx)
	user, err := a.authRepo.Login(tx, request)
	if err != nil {
		if err == constants.ErrUserNotFound {
			a.log.Info("find user by email")
			return model.ResponseInterface{
				Error:      err,
				StatusCode: http.StatusBadRequest,
				Message:    constants.InvalidEmailAndPassword,
			}
		}
		a.log.Info("find user by email", zap.Error(err))
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusBadRequest,
			Message:    constants.InvalidEmailAndPassword,
		}
	}

	isValid := util.CheckPasswordHash(request.Password, user.Password)

	if !isValid {
		a.log.Info("checking password", zap.Error(constants.ErrInvalidUsernameAndPassword))

		return model.ResponseInterface{
			Error:      constants.ErrInvalidUsernameAndPassword,
			StatusCode: http.StatusBadRequest,
			Message:    constants.InvalidEmailAndPassword,
		}
	}

	jwtToken, expiredAt, err := util.GenerateJwtToken(util.JWTParams{
		Payload:    user.UUID,
		SecretKey:  a.config.JWT_SECRET_KEY,
		ExpireTime: a.config.JWT_EXPIRE_TIME,
	})

	if err != nil {
		a.log.Info("Generate token", zap.Error(err))
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	jwtRefreshToken, _, err := util.GenerateJwtToken(util.JWTParams{
		Payload:    user.UUID,
		SecretKey:  a.config.JWT_SECRET_KEY,
		ExpireTime: a.config.JWT_EXPIRE_TIME,
	})

	if err != nil {
		a.log.Info("Generate token refresh token", zap.Error(err))

		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	err = a.insertTokenToRedis(ctx, user, jwtToken, jwtRefreshToken)

	if err != nil {
		a.log.Info("insert redis", zap.Error(err))

		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	return model.ResponseInterface{
		StatusCode: http.StatusOK,
		Data: model.LoginResponse{
			Token:        jwtToken,
			RefreshToken: jwtRefreshToken,
			ExpiredAt:    expiredAt.String(),
		},
		Message: constants.Success,
	}
}

func (a *AuthUsecase) Register(ctx *gin.Context, request model.RegisterRequest) model.ResponseInterface {
	tx := a.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := a.userRepo.CountByEmail(tx, request.Email)
	if err == nil && user != 0 {
		a.log.Info("when checking email", zap.Error(err))
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusBadRequest,
			Message:    constants.EmailAlreadyExists,
		}
	}

	hashPassword, err := util.HashPassword(request.Password)
	if err != nil {
		a.log.Info("hashing password", zap.Error(err))
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}

	var newUser entity.User
	newUser.Email = request.Email
	newUser.UserCode = util.GenerateNumber(4)
	newUser.Password = hashPassword

	newUser.PhoneNumber = request.PhoneNumber

	err = a.userRepo.Create(tx, &newUser)
	if err != nil {
		a.log.Info("create users", zap.Error(err))
		return model.ResponseInterface{
			Error:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    constants.InternalServerError,
		}
	}
	if err := tx.Commit().Error; err != nil {
		a.log.Info("commit transaction", zap.Error(err))
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

func (a *AuthUsecase) insertTokenToRedis(ctx *gin.Context, user *entity.User, jwtToken string, jwtRefreshToken string) error {
	keyToken := fmt.Sprintf("%s_%s", user.UUID.String(), constants.Token)
	_, err := a.redis.Set(ctx, keyToken, jwtToken, time.Duration(a.config.JWT_EXPIRE_TIME)*time.Minute).Result()

	if err != nil {
		return err
	}

	keyRefreshToken := fmt.Sprintf("%s_%s", user.UUID.String(), constants.RefreshToken)
	_, err = a.redis.Set(ctx, keyRefreshToken, jwtRefreshToken, time.Duration(a.config.JWT_EXPIRE_REFRESH_TIME)*time.Minute).Result()

	if err != nil {
		return err
	}
	return nil

}

func NewAuthUsecase(
	db *gorm.DB,
	log *zap.Logger,
	config util.Config,
	redis *redis.Client,
	authRepo *repository.AuthRepository,
	userRepo *repository.UserRepository,
) *AuthUsecase {
	return &AuthUsecase{
		db:       db,
		log:      log,
		config:   config,
		redis:    redis,
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
