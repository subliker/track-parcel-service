package session

import "errors"

var (
	// error if session is already exist
	ErrSessionIsAlreadyExist = errors.New("session is already exist")
	// error if session is not exist
	ErrSessionIsNotExist = errors.New("session is not exist")

	// error if state wasn't done but required
	ErrStateNotDone = errors.New("state isn't done")
	// error if response ended with not found
	ErrResNotFound = errors.New("response error: not found")
	// error if not specify usage was incorrect
	ErrIncorrectNotSpecify = errors.New("not specify is called in not current field")
)
