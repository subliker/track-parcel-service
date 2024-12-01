package parcel

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

type NotificationStore interface {
	ParcelSubscribers(model.TrackNumber) ([]model.TelegramID, error)
	Close()
}
