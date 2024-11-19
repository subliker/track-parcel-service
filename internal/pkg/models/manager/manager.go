package manager

import "github.com/subliker/track-parcel-service/internal/pkg/models/telegram"

type ApiToken string
type Manager struct {
	TelegramId  telegram.ID
	FullName    string
	Email       string
	PhoneNumber *string
	Company     *string
}
