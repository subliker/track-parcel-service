package model

type User struct {
	TelegramID  TelegramID `validate:"required"`
	FullName    string     `validate:"required,min=3,max=255"`
	Email       string     `validate:"required,email"`
	PhoneNumber *string    `validate:"omitempty,e164"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
