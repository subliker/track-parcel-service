package user

import "errors"

var (
	ErrUserIsAlreadyExist = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user is not found")
	ErrInternal              = errors.New("internal server error")
	ErrUnexpected            = errors.New("unexpected error")
)
