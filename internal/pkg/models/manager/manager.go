package manager

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
)

type ApiToken string

func NewApiToken() (ApiToken, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(bytes)
	return ApiToken(token), nil
}

type Manager struct {
	TelegramId  telegram.ID
	FullName    string
	Email       string
	PhoneNumber *string
	Company     *string
}
