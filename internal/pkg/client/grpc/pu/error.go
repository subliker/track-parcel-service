package pu

import "errors"

var (
	ErrParcelNotFound       = errors.New("parcel is not found")
	ErrSubscriptionNotFound = errors.New("subscription is not found")
	ErrInternal             = errors.New("internal server error")
	ErrUnexpected           = errors.New("unexpected error")
	ErrAlreadyExists        = errors.New("already exists")
)
