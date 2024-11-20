package store

import "errors"

var (
	ErrUserNotFound = errors.New("user wasn't found")

	ErrManagerNotFound = errors.New("manager wasn't found")
)
