package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
)

type AuthController struct {
	log         *zap.Logger
	authUsecase *usecase.AuthUsecase
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body model.LoginRequest true "Login Request"
// @Success 200 {object} model.Response{response_data=model.LoginResponse}
// @Router /auth/login [post]
func (a *AuthController) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		a.log.Info("error binding request", zap.Error(err))

		util.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	response := a.authUsecase.Login(ctx, loginRequest)
	util.Response(ctx, response)
}

// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param registerRequest body model.RegisterRequest true "Register Request"
// @Success 200 {object} model.Response{response_data=string}
// @Router /auth/register [post]
func (a *AuthController) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		a.log.Info("error binding request", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	response := a.authUsecase.Register(ctx, registerRequest)
	util.Response(ctx, response)
}

func NewAuthController(authUsecase *usecase.AuthUsecase, log *zap.Logger) *AuthController {
	return &AuthController{
		log:         log,
		authUsecase: authUsecase,
	}
}
