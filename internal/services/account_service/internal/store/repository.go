package store

import (
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
)

//go:generate mockgen -source=repository.go -destination=mock/repository.go -package=mock

// UserRepository is interface for accessing user repository
type UserRepository interface {
	// Register adds user to repository
	Register(model.User) error
	// Get returns user's info from repository
	Get(model.TelegramID) (model.User, error)
	// Exists returns true if user with telegram id exists in repository
	Exists(model.TelegramID) (bool, error)
}

// ManagerRepository is interface for accessing manager repository
type ManagerRepository interface {
	// Register adds manager to repository
	Register(model.Manager) error
	// Get returns manager's with telegram id info from repository
	Get(model.TelegramID) (model.Manager, error)
	// RetrieveManagerIdByApiKey returns api token owner's telegram id from repository
	RetrieveManagerIdByApiKey(apiToken model.ManagerApiToken) (model.TelegramID, error)
	// GetApiToken returns manager's with telegram id api token from repository
	GetApiToken(model.TelegramID) (model.ManagerApiToken, error)
	// Exists returns true if manager with telegram id exists in repository
	Exists(model.TelegramID) (bool, error)
}
