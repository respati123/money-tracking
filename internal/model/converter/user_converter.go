package converter

import (
	"time"

	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/model"
)

type UserConverter struct {
	// ctx *gin.Context
}

func NewUserConverter() UserConverter {
	return UserConverter{
		// ctx: ctx,
	}
}

func (u *UserConverter) ToResponseUser(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:          user.ID,
		UserCode:    user.UserCode,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
		CreatedBy:   user.CreatedBy,
		UpdatedBy:   user.UpdatedBy,
	}
}

func (u *UserConverter) ToResponseUsers(users *[]entity.User) *[]model.UserResponse {
	var userResponses []model.UserResponse
	for _, user := range *users {
		userResponses = append(userResponses, *u.ToResponseUser(&user))
	}

	return &userResponses
}
