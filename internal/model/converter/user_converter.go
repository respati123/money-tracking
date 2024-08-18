package converter

import (
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
)

type UserConverter struct{}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (uc *UserConverter) ToUserResponse(user *entity.User) *model.UserResponse {

	return &model.UserResponse{
		ID:          user.ID,
		UUID:        user.UUID.String(),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UserCode:    user.UserCode,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   &user.DeletedAt.Time,
		CreatedBy:   user.CreatedBy,
		UpdatedBy:   user.UpdatedBy,
		DeletedBy:   user.DeletedBy,
	}
}

func (uc *UserConverter) ToUserResponses(users *[]entity.User) *[]model.UserResponse {
	var userResponse []model.UserResponse
	for _, user := range *users {
		userResponse = append(userResponse, *uc.ToUserResponse(&user))
	}
	return &userResponse
}
