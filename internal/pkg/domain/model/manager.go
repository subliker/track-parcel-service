package model

import (
	"crypto/rand"
	"encoding/base64"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
	"regexp"
)

func init() {
	logger := zap.Logger.WithFields("layer", "domain")

	if err := validator.V.RegisterValidation("fullName", fullNameValidator); err != nil {
		logger.Fatal(err)
	}
}

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
	FullName    string     `validate:"required,min=3,max=255,fullName"`
	Email       string     `validate:"required,email"`
	PhoneNumber *string    `validate:"omitempty,e164"`
	Company     *string    `validate:"omitempty,min=3,max=255"`
}

var _fullNameRe = regexp.MustCompile(`^[а-яА-ЯёЁa-zA-Z]+ [а-яА-ЯёЁa-zA-Z]+ ?[а-яА-ЯёЁa-zA-Z]+$`)

func fullNameValidator(fl validator2.FieldLevel) bool {
	return _fullNameRe.MatchString(fl.Field().String())
}

func (m *Manager) Validate() error {
	return validator.V.Struct(m)
}

// FormatToE164 removes all extra symbols from phone number string.
func FormatToE164(phone string) string {
	return regexp.MustCompile(`[^\d+]`).ReplaceAllString(phone, "")
}
