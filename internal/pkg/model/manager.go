package model

import (
	"crypto/rand"
	"encoding/base64"
)

type ManagerApiToken string

func NewManagerApiToken() (ManagerApiToken, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(bytes)
	return ManagerApiToken(token), nil
}

type Manager struct {
	TelegramID  TelegramID
	FullName    string
	Email       string
	PhoneNumber *string
	Company     *string
}
