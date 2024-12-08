package parcel

import "github.com/subliker/track-parcel-service/internal/pkg/model"

type UserStore interface {
	GetInfo(model.TrackNumber) (model.Parcel, error)
	GetSubscribed(model.TrackNumber, model.TelegramID) (bool, error)
	GetCheckpoints(trackNum model.TrackNumber, page uint64, pageSize uint64) ([]*model.Checkpoint, error)
	AddSubscription(model.TrackNumber, model.TelegramID) error
	DeleteSubscription(model.TrackNumber, model.TelegramID) error
	Exists(model.TrackNumber) (bool, error)
	Close()
}
