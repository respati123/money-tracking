package http

import (
	"net/http"

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
		util.SendErrorResponse(ctx, http.StatusBadRequest, "error while error")
		return
	}

	response := rc.roleUsecase.Create(ctx, request)
	util.Response(ctx, response)
}
