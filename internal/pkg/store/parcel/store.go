package parcel

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

// Store is interface for accessing parcels store
//
//go:generate mockgen -source=store.go -destination=mock/store.go -package=mock
type Store interface {
	// CheckAccess checks whether the package belongs to the manager
	CheckAccess(model.TrackNumber, model.TelegramID) (bool, error)
	// Add adds parcel in store
	Add(model.Parcel) (model.TrackNumber, error)
	// Delete deletes parcel from store by track number
	Delete(model.TrackNumber) error
	// GetInfo returns parcel info from store by track number
	GetInfo(model.TrackNumber) (model.Parcel, error)
	// GetSubscribed returns true if user with telegram id if subscribed on parcel with track number
	GetSubscribed(model.TrackNumber, model.TelegramID) (bool, error)
	// Exists checks if parcel with track number exists in store
	Exists(model.TrackNumber) (bool, error)

	// AddCheckpoint adds checkpoint in store for parcel with track number
	AddCheckpoint(model.TrackNumber, model.Checkpoint) error
	// GetCheckpoints returns parcel's with track number checkpoints from store with pagination
	GetCheckpoints(trackNum model.TrackNumber, page uint64, pageSize uint64) ([]*model.Checkpoint, error)

	// AddSubscription adds user's with telegram id subscription for parcel with track number in store
	AddSubscription(model.TrackNumber, model.TelegramID) error
	// DeleteSubscription removes user's with telegram id subscription for parcel with track number from store
	DeleteSubscription(model.TrackNumber, model.TelegramID) error
	// ParcelSubscribers returns array of user's telegram id subscribed for parcel with track number
	ParcelSubscribers(model.TrackNumber) ([]model.TelegramID, error)

	// Close closes store
	Close() error
}
