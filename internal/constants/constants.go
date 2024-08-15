package constants

import "errors"

var (
	// REsponse JSON
	Response = "response"
	UserData = "userdata"

	// Message
	Message = "message"

	// Error
	ErrUserAlreadyExist           = errors.New("user_already_exists")
	ErrUserNotFound               = errors.New("user_not_found")
	ErrEmailAlreadyExist          = errors.New("email_already_exists")
	ErrWrongPassword              = errors.New("wrong_password")
	ErrInvalidUsernameAndPassword = errors.New("invalid_username_and_password")

	// // Response Message

	//Error Message
	UserNotFound            = "User Not Found"
	EmailAlreadyExists      = "Email Already Exists"
	UserAlreadyExists       = "User Already Exists"
	InvalidEmailAndPassword = "Invalid username and password"
	InternalServerError     = "Internal server error"

	// Success Message
	Success = "Succeed"
	Error   = "Error"

	// strings
	Token        = "token"
	RefreshToken = "refreshtoken"
)
