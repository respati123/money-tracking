package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/configs/logger"
	"github.com/respati123/money-tracking/internal/model"
	"github.com/respati123/money-tracking/internal/usecase"
	"github.com/respati123/money-tracking/internal/util"
)

type UserController struct {
	userUseCase *usecase.UserUseCase
	log         *logger.CustomLogger
}

func NewUserController(userUseCase *usecase.UserUseCase, log *logger.CustomLogger) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		log:         log.Module("user-controller"),
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
		a.log.ErrorWithFields(ctx, "error binding request ", err)
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error binding request", err)
		return
	}
	response, err := a.userUseCase.GetListUser(ctx, request)

	if err != nil {
		a.log.ErrorWithFields(ctx, "error get list user ", err)
		util.SendErrorResponse(ctx, http.StatusInternalServerError, "error get list user", err)
		return
	}

	util.SendSuccessResponse(ctx, http.StatusOK, "success get list user", response)

}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Produce json
// @Tags Users
// @Param body body model.UserCreateRequest true "Body User Create"
// @Success 200 {object} model.Response{response_data=string} "success create user"
// @Router /users/create [post]
// @Security ApiKeyAuth
func (a *UserController) CreateUser(ctx *gin.Context) {
	var (
		request model.UserCreateRequest
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		a.log.ErrorWithFields(ctx, "error binding request ", err)
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error binding request", err)
		return
	}

	err := a.userUseCase.CreateUser(ctx, request)

	if err != nil {
		a.log.ErrorWithFields(ctx, "error create user ", err)
		util.SendErrorResponse(ctx, http.StatusInternalServerError, "internal server error", err)
		return
	}
	util.SendSuccessResponse(ctx, http.StatusOK, "success create user", nil)
}

// Detele user
// @Summary Delete user
// @Description Delete a new user
// @Produce json
// @Tags Users
// @Param uuid path string true  "uuid user"
// @Success 200 {object} model.Response{response_data=string} "success delete user"
// @Router /users/delete/{uuid} [delete]
// @Security ApiKeyAuth
func (a *UserController) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := a.userUseCase.DeleteUser(ctx, uuid)
	if err != nil {
		a.log.ErrorWithFields(ctx, "error delete user", err)
		util.SendErrorResponse(ctx, http.StatusInternalServerError, "internal server error", err)
		return
	}

	util.SendSuccessResponse(ctx, http.StatusOK, "success delete user", nil)
}
