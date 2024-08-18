package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
	"go.uber.org/zap"
)

type UserController struct {
	userUseCase *usecase.UserUseCase
	log         *zap.Logger
}

func NewUserController(userUseCase *usecase.UserUseCase, log *zap.Logger) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		log:         log,
	}
}

// GetListUser retrieves a list of users.
// @Summary Get a list of users
// @Description Retrieves a list of users
// @Produce json
// @Tags Users
// @Param body body model.PaginationRequest false "Body Pagination"
// @Success 200 {object} model.Response{response_data=model.PaginationResponse{data=model.UserResponse}} "success get list user"
// @Router /users/list [post]
// @Security ApiKeyAuth
func (a *UserController) GetListUser(ctx *gin.Context) {
	var (
		request model.PaginationRequest
	)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		a.log.Info("error binding request", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error binding request", err)
		return
	}
	response := a.userUseCase.GetListUser(ctx, request)
	util.Response(ctx, response)
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Produce json
// @Tags Users
// @Param body body model.UserCreateRequest true "Body User Create"
// @Success 200 {object} model.Response{response_data=string} "success create user"
// @Router /users/ [post]
// @Security ApiKeyAuth
func (a *UserController) CreateUser(ctx *gin.Context) {
	var (
		request model.UserCreateRequest
	)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		a.log.Info("error binding request", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error binding request", err)
		return
	}
	response := a.userUseCase.CreateUser(ctx, request)
	util.Response(ctx, response)
}

// Detele user
// @Summary Delete user
// @Description Delete a new user
// @Produce json
// @Tags Users
// @Param uuid path string true  "uuid user"
// @Success 200 {object} model.Response{response_data=string} "success delete user"
// @Router /users/{uuid} [delete]
// @Security ApiKeyAuth
func (a *UserController) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	response := a.userUseCase.DeleteUser(ctx, uuid)
	util.Response(ctx, response)
}

// Update User
// @Summary Update User
// @Description Update User
// @Produce json
// @Tags Users
// @Param uuid path string true  "uuid user"
// @Success 200 {object} model.Response{response_data=string} "success update user"
// @Router /users/{uuid} [put]
// @Security ApiKeyAuth
func (a *UserController) Update(ctx *gin.Context) {
	var (
		request model.UserUpdateRequest
	)
	id := ctx.Param("uuid")

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		a.log.Info("error binding request", zap.Error(err))
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error binding request", err)
		return
	}

	a.userUseCase.UpdateUser(ctx, id, request)
}

// Get User
// @Summary Get User
// @Description Get User
// @Produce json
// @Tags Users
// @Param user_code path int true  "user_code user"
// @Success 200 {object} model.Response{response_data=model.UserResponse} "success Get user"
// @Router /users/{user_code} [get]
// @Security ApiKeyAuth
func (a *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("user_code")
	code, _ := strconv.Atoi(id)
	response := a.userUseCase.GetUser(ctx, code)

	util.Response(ctx, response)
}
