package store

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

type UserRepository interface {
	Register(model.User) error
	Get(model.TelegramID) (model.User, error)
	Exists(model.TelegramID) (bool, error)
}

type ManagerRepository interface {
	Register(model.Manager) error
	Get(model.TelegramID) (model.Manager, error)
	GetTelegramId(apiToken model.ManagerApiToken) (model.TelegramID, error)
	GetApiToken(model.TelegramID) (model.ManagerApiToken, error)
	Exists(model.TelegramID) (bool, error)
}
