package model

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/subliker/track-parcel-service/internal/pkg/validator"
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
	TelegramID  TelegramID `validate:"required"`
	FullName    string     `validate:"required,min=3,max=255"`
	Email       string     `validate:"required,email"`
	PhoneNumber *string    `validate:"omitempty,e164"`
	Company     *string    `validate:"omitempty,min=3,max=255"`
}

func (m *Manager) Validate() error {
	return validator.V.Struct(m)
}
