package parcel

import (
	model2 "github.com/subliker/track-parcel-service/internal/pkg/domain/model"
)

// Store is interface for accessing parcels store
//
//go:generate mockgen -source=store.go -destination=mock/store.go -package=mock
type Store interface {
	// CheckAccess checks whether the package belongs to the manager
	CheckAccess(model2.TrackNumber, model2.TelegramID) (bool, error)
	// Add adds parcel in store
	Add(model2.Parcel) (model2.TrackNumber, error)
	// Delete deletes parcel from store by track number
	Delete(model2.TrackNumber) error
	// GetInfo returns parcel info from store by track number
	GetInfo(model2.TrackNumber) (model2.Parcel, error)
	// GetSubscribed returns true if user with telegram id if subscribed on parcel with track number
	GetSubscribed(model2.TrackNumber, model2.TelegramID) (bool, error)
	// Exists checks if parcel with track number exists in store
	Exists(model2.TrackNumber) (bool, error)

	// AddCheckpoint adds checkpoint in store for parcel with track number
	AddCheckpoint(model2.TrackNumber, model2.Checkpoint) error
	// GetCheckpoints returns parcel's with track number checkpoints from store with pagination
	GetCheckpoints(trackNum model2.TrackNumber, page uint64, pageSize uint64) ([]*model2.Checkpoint, error)

	// AddSubscription adds user's with telegram id subscription for parcel with track number in store
	AddSubscription(model2.TrackNumber, model2.TelegramID) error
	// DeleteSubscription removes user's with telegram id subscription for parcel with track number from store
	DeleteSubscription(model2.TrackNumber, model2.TelegramID) error
	// ParcelSubscribers returns array of user's telegram id subscribed for parcel with track number
	ParcelSubscribers(model2.TrackNumber) ([]model2.TelegramID, error)

	// Close closes store
	Close() error
}
