package manager

import "errors"

var (
	ErrManagerIsAlreadyExist = errors.New("manager already exists")
	ErrManagerNotFound       = errors.New("manager is not found")
	ErrInternal              = errors.New("internal server error")
	ErrUnexpected            = errors.New("unexpected error")
)
