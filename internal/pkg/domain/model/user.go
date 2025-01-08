package model

import "github.com/subliker/track-parcel-service/internal/pkg/validator"

type User struct {
	TelegramID  TelegramID `validate:"required"`
	FullName    string     `validate:"required,min=3,max=255,fullName"`
	Email       string     `validate:"required,email"`
	PhoneNumber *string    `validate:"omitempty,e164"`
}

func (u *User) Validate() error {
	return validator.V.Struct(u)
}
