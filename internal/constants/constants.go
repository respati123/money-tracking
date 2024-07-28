package constants

import "errors"

var (
	// REsponse JSON
	Response = "response"

	// Message
	Message = "message"

	// Error
	ErrUserAlreadyExist = errors.New("user_already_exists")
	ErrUserNotFound     = errors.New("user_not_found")
)
