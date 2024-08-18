package constants

import "errors"

func ErrNotFound(text string) error {
	return errors.New(text + " not found")
}
func ErrDuplicate(text string) error {
	return errors.New(text + " is duplicate")
}

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
	ErrInternalServerError        = errors.New("internal_server_error")

	// // Response Message

	//Error Message
	UserNotFound            = "User Not Found"
	EmailAlreadyExists      = "Email Already Exists"
	UserAlreadyExists       = "User Already Exists"
	InvalidEmailAndPassword = "Invalid username and password"
	InternalServerError     = "Internal server error"
	DuplicateKey            = "duplicate key value violates unique constraint"

	// Success Message
	Success = "Succeed"
	Error   = "Error"

	// strings
	Token        = "token"
	RefreshToken = "refreshtoken"
	TraceID      = "trace_id"
)
