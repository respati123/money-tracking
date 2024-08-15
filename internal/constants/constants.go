package constants

import "errors"

var (
	// REsponse JSON
	Response = "response"
	UserData = "userdata"

	// Message
	Message = "message"

	// Context
	TraceID = "trace_id"

	// Error
	ErrUserAlreadyExist  = errors.New("user_already_exists")
	ErrUserNotFound      = errors.New("user_not_found")
	ErrEmailAlreadyExist = errors.New("email_already_exists")
)
