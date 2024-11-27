package state

import "errors"

var (
	ErrStateNotDone        = errors.New("state isn't done")
	ErrResNotFound         = errors.New("response error: not found")
	ErrIncorrectNotSpecify = errors.New("not specify is called in not current field")
)
