package store

import (
	"github.com/subliker/track-parcel-service/internal/pkg/models/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/models/user"
)

type UserRepository interface {
	Register(user.User) error
	Get(telegram.ID) (user.User, error)
}

type ManagerRepository interface {
	Register(manager.Manager) error
	Get(telegram.ID) (manager.Manager, error)
	GetApiToken(telegram.ID) (manager.ApiToken, error)
}
