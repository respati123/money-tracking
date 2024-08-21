package model

import "time"

type UserFilter struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
type UserCreateRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserResponse struct {
	ID          uint       `json:"id"`
	UUID        string     `json:"uuid"`
	UserCode    int        `json:"user_code"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	CreatedBy   int        `json:"created_by"`
	UpdatedBy   int        `json:"updated_by"`
	DeletedBy   int        `json:"deleted_by"`
}

type UserUpdateRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserUpdatePasswordRequest struct {
	Password string `json:"password"`
}
