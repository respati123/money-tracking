package converter

import (
	"time"

	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
)

type RoleConverter struct{}

func NewRoleConverter() *RoleConverter {
	return &RoleConverter{}
}

func (nr *RoleConverter) ToRoleResponse(role *entity.Role) *model.RoleResponse {

	return &model.RoleResponse{
		ID:        int(role.ID),
		UUID:      role.UUID,
		Name:      role.Name,
		Alias:     role.Alias,
		CreatedAt: role.CreatedAt.Format(time.RFC3339),
		UpdatedAt: role.UpdatedAt.Format(time.RFC3339),
		DeletedAt: role.DeletedAt.Time.Format(time.RFC3339),
		CreatedBy: role.CreatedBy,
		UpdatedBy: role.UpdatedBy,
		DeletedBy: role.DeletedBy,
	}
}

func (nr *RoleConverter) ToRoleResponses(roles *[]entity.Role) *[]model.RoleResponse {
	var roleResponse []model.RoleResponse
	for _, role := range *roles {
		roleResponse = append(roleResponse, *nr.ToRoleResponse(&role))
	}
	return &roleResponse
}
