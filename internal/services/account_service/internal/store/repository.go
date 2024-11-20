package store

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

type UserRepository interface {
	Register(model.User) error
	Get(model.TelegramID) (model.User, error)
}

type ManagerRepository interface {
	Register(model.Manager) error
	Get(model.TelegramID) (model.Manager, error)
	GetApiToken(model.TelegramID) (model.ManagerApiToken, error)
}
