package parcel

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

// NotificationStore stores data about subscriptions
type NotificationStore interface {
	// ParcelSubscribers returns array of user telegram id subscribers
	ParcelSubscribers(model.TrackNumber) ([]model.TelegramID, error)
	// Close closes store
	Close()
}
