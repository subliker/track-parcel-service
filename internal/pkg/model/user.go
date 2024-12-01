package model

type User struct {
	TelegramID  TelegramID
	FullName    string
	Email       string
	PhoneNumber *string
}
