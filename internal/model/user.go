package model

type UserCreateRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserResponse struct {
	ID          uint   `json:"id"`
	UUID        string `json:"uuid"`
	UserCode    int    `json:"user_code"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}

type UserUpdateRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserUpdatePasswordRequest struct {
	Password string `json:"password"`
}
