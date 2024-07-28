package usecase

import (
	"context"

	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/util"
	"github.com/sirupsen/logrus"
)

type AuthUsecase interface {
	Login(ctx context.Context, request model.LoginRequest) (model.LoginResponse, error)
	Register(ctx context.Context, request model.RegisterRequest) error
}

type authUsecase struct {
	config   util.Config
	log      *logrus.Logger
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
}

func (a *authUsecase) Login(ctx context.Context, request model.LoginRequest) (model.LoginResponse, error) {

	user, err := a.authRepo.Login(ctx, request)
	if err != nil {
		if err == constants.ErrUserNotFound {
			a.log.WithFields(logrus.Fields{
				"trace_id": ctx.Value("trace_id"),
				"module":   "auth_usecase",
				"method":   "Login",
			}).Error("password or email invalid :", err)
			return model.LoginResponse{}, err
		}
		return model.LoginResponse{}, err
	}

	isValid := util.CheckPasswordHash(request.Password, user.Password)

	if !isValid {
		a.log.WithFields(logrus.Fields{
			"trace_id": ctx.Value("trace_id"),
			"module":   "auth_usecase",
			"method":   "Login",
		}).Error("password or email invalid :", err)
		return model.LoginResponse{}, err
	}

	jwtToken, expiredAt, err := util.GenerateJwtToken(util.JWTParams{
		Payload:    user.UUID,
		SecretKey:  a.config.SECRET_KEY_JWT,
		ExpireTime: a.config.JWT_EXPIRE_TIME,
	})

	if err != nil {
		a.log.WithFields(logrus.Fields{
			"trace_id": ctx.Value("trace_id"),
			"module":   "auth_usecase",
			"method":   "Login",
		}).Error("error generate jwt token :", err)
		return model.LoginResponse{}, err
	}

	jwtRefreshToken, _, err := util.GenerateJwtToken(util.JWTParams{
		Payload:    user.UUID,
		SecretKey:  a.config.SECRET_KEY_JWT,
		ExpireTime: a.config.JWT_REFRESH_EXPIRE_TIME,
	})

	if err != nil {
		a.log.WithFields(logrus.Fields{
			"trace_id": ctx.Value("trace_id"),
			"module":   "auth_usecase",
			"method":   "Login",
		}).Error("error generate jwt refresh token :", err)
		return model.LoginResponse{}, err
	}

	a.log.WithFields(logrus.Fields{
		"trace_id": ctx.Value("trace_id"),
		"module":   "auth_usecase",
		"method":   "Login",
	}).Info("login successfully")

	return model.LoginResponse{
		Token:        jwtToken,
		RefreshToken: jwtRefreshToken,
		ExpiredAt:    expiredAt.String(),
	}, nil
}

func (a *authUsecase) Register(ctx context.Context, request model.RegisterRequest) error {
	a.log.WithFields(logrus.Fields{
		"trace_id": ctx.Value("trace_id"),
		"module":   "auth_usecase",
		"method":   "Register",
	}).Info("register request ")

	user, err := a.userRepo.GetUserIDByEmail(ctx, request.Email)
	if err == nil && user != 0 {
		a.log.WithFields(logrus.Fields{
			"trace_id": ctx.Value("trace_id"),
			"module":   "auth_usecase",
			"method":   "Register",
		}).Error("email already exist")
		return constants.ErrUserAlreadyExist
	}

	hashPassword, err := util.HashPassword(request.Password)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"trace_id": ctx.Value("trace_id"),
			"module":   "auth_usecase",
			"method":   "Register",
		}).Error("error hashing password")
		return err
	}

	var newUser entity.User
	newUser.Email = request.Email
	newUser.UserCode = util.GenerateNumber(4)
	newUser.Password = hashPassword
	newUser.PhoneNumber = request.PhoneNumber

	_, err = a.userRepo.CreateUser(ctx, &newUser)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"trace_id": ctx.Value("trace_id"),
			"module":   "auth_usecase",
			"method":   "Register",
		}).Error("error create user", err)
		return err
	}

	a.log.WithFields(logrus.Fields{
		"trace_id": ctx.Value("trace_id"),
		"module":   "auth_usecase",
		"method":   "Register",
	}).Info("register successfully")

	return nil

}

func NewAuthUsecase(log *logrus.Logger, authRepo repository.AuthRepository, userRepo repository.UserRepository, config util.Config) AuthUsecase {
	return &authUsecase{
		log:      log,
		authRepo: authRepo,
		userRepo: userRepo,
		config:   config,
	}
}
