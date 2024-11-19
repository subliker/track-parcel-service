package session

import "errors"

var (
	// error if session is already exist
	ErrSessionIsAlreadyExist = errors.New("session is already exist")
	// error if session is not exist
	ErrSessionIsNotExist = errors.New("session is not exist")
)
