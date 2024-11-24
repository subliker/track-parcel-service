package pm

import "errors"

var (
	ErrParcelNotFound = errors.New("parcel is not found")
	ErrInternal       = errors.New("internal server error")
	ErrUnexpected     = errors.New("unexpected error")
)
