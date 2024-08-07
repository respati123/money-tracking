package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
)

type AuthController struct {
	log         *logger.CustomLogger
	authUsecase *usecase.AuthUsecase
}

// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body model.LoginRequest true "Login Request"
// @Success 200 {object} model.Response{response_data=model.LoginResponse}
// @Router /auth/login [post]
func (a *AuthController) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		a.log.ErrorWithFields(ctx, "error invalid request", err)

		util.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	response, err := a.authUsecase.Login(ctx, loginRequest)
	if err != nil {
		if err == constants.ErrUserNotFound {
			a.log.ErrorWithFields(ctx, "error user not found", err)

			util.SendErrorResponse(ctx, http.StatusNotFound, "user not found", err)
			return
		}
		a.log.ErrorWithFields(ctx, "internal server error", err)
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
func (a *AuthController) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		a.log.ErrorWithFields(ctx, "error binding request", err)
		util.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	err = a.authUsecase.Register(ctx, registerRequest)
	if err != nil {
		if err == constants.ErrUserAlreadyExist {
			a.log.ErrorWithFields(ctx, "error email already exist", err)
			util.SendErrorResponse(ctx, http.StatusBadRequest, "email already exist", err)
			return
		}
		a.log.ErrorWithFields(ctx, "error bad request register", err)

		util.SendErrorResponse(ctx, http.StatusBadRequest, "internal server error", err)
		return
	}
	util.SendSuccessResponse(ctx, http.StatusCreated, "register success", nil)
}

func NewAuthController(authUsecase *usecase.AuthUsecase, log *logger.CustomLogger) *AuthController {
	return &AuthController{
		log:         log.Module("auth-controller"),
		authUsecase: authUsecase,
	}
}
