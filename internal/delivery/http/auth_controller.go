package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"github.com/sirupsen/logrus"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	log         *logrus.Logger
	authUsecase usecase.AuthUsecase
}

// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body model.LoginRequest true "Login Request"
// @Success 200 {object} model.Response{response_data=model.LoginResponse}
// @Router /auth/login [post]
func (a *authController) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		util.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	response, err := a.authUsecase.Login(ctx, loginRequest)
	if err != nil {
		if err == constants.ErrUserNotFound {
			util.SendErrorResponse(ctx, http.StatusNotFound, "user not found", err)
			return
		}
		util.SendErrorResponse(ctx, http.StatusInternalServerError, "internal server error", err)
		return
	}
	util.SendSuccessResponse(ctx, http.StatusOK, "login success", response)
}

// @Summary Register
// @Description Register
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body model.RegisterRequest true "Register Request"
// @Success 200 {object} model.Response{response_data=string}
// @Router /auth/register [post]
func (a *authController) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"module": "auth_controller",
			"method": "Register",
		}).Error("error bind json :", err)
		util.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	err = a.authUsecase.Register(ctx, registerRequest)

	if err != nil {
		if err == constants.ErrUserAlreadyExist {
			util.SendErrorResponse(ctx, http.StatusBadRequest, "email already exist", err)
			return
		}
		a.log.WithFields(logrus.Fields{
			"module": "auth_controller",
			"method": "Register",
		}).Error("Bad Request ", err)
		util.SendErrorResponse(ctx, http.StatusBadRequest, "internal server error", err)
		return
	}

	util.SendSuccessResponse(ctx, http.StatusCreated, "register success", nil)
}

func NewAuthController(log *logrus.Logger, authUsecase usecase.AuthUsecase) AuthController {
	return &authController{
		log:         log,
		authUsecase: authUsecase,
	}
}
