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

type RoleController struct {
	roleUsecase *usecase.RoleUsecase
	log         *zap.Logger
}

func NewRoleController(roleUsecase *usecase.RoleUsecase, log *zap.Logger) *RoleController {
	return &RoleController{
		log:         log,
		roleUsecase: roleUsecase,
	}
}

// Create Role
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Produce json
// @Accept json
// @Param body body model.RoleCreateRequest true "Role Create Request"
// @Success 200 {object} model.Response{response_data=string} "Successfully create role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /roles/ [post]
// @Security ApiKeyAuth
func (rc *RoleController) Create(ctx *gin.Context) {
	var (
		request model.RoleCreateRequest
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while error", err)
		return
	}

	response := rc.roleUsecase.Create(ctx, request)
	util.Response(ctx, response)
}

// Update Role
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Produce json
// @Accept json
// @Param uuid path string true "uuid"
// @Param body body model.RoleUpdateRequest true "Role update request"
// @Success 200 {object} model.Response{response_data=string} "successfully update role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /roles/{uuid} [put]
// @Security ApiKeyAuth
func (rc *RoleController) Update(ctx *gin.Context) {

	var (
		request model.RoleUpdateRequest
	)
	paramId := ctx.Param("uuid")

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while binding", err)
		return
	}

	response := rc.roleUsecase.Update(ctx, request, paramId)
	util.Response(ctx, response)
}

// Delete Role
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Produce json
// @Accept json
// @Param uuid path string true "role uuid"
// @Success 200 {object} model.Response{response_data=string} "successfully delete role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /roles/{uuid} [delete]
// @Security ApiKeyAuth
func (rc *RoleController) Delete(ctx *gin.Context) {
	paramId := ctx.Param("uuid")

	response := rc.roleUsecase.Delete(ctx, paramId)
	util.Response(ctx, response)
}

// Find all Role
// @Summary Find all Role
// @Description Find all Role
// @Tags Role
// @Produce json
// @Accept json
// @Param body body model.PaginationRequest true "pagination request"
// @Success 200 {object} model.Response{response_data=string} "successfully find all role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /roles/all [post]
// @Security ApiKeyAuth
func (rc *RoleController) FindAll(ctx *gin.Context) {
	var (
		request model.PaginationRequest
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while binding", err)
		return
	}

	response := rc.roleUsecase.FindAll(ctx, request)
	util.Response(ctx, response)
}

// Get Role
// @Summary Get Role
// @Description Get Role
// @Tags Role
// @Produce json
// @Accept json
// @Param role_code path int true "role code"
// @Success 200 {object} model.Response{response_data=model.RoleResponse} "successfully get role"
// @Failure 400 {object} model.Response{response_data=string} "Bad Request"
// @Failure 500 {object} model.Response{response_data=string} "Internal Server Error"
// @Router /roles/{role_code} [get]
// @Security ApiKeyAuth
func (rc *RoleController) GetRole(ctx *gin.Context) {
	paramId := ctx.Param("role_code")

	code, err := strconv.Atoi(paramId)

	if err != nil {
		rc.log.Error("error convert paramId")
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error convert paramId", err)
		return
	}

	response := rc.roleUsecase.GetRole(ctx, code)
	util.Response(ctx, response)
}
