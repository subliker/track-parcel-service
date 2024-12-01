package parcel

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

type Store interface {
	CheckAccess(model.TrackNumber, model.TelegramID) (bool, error)
	Add(model.Parcel) (model.TrackNumber, error)
	Delete(model.TrackNumber) error
	GetInfo(model.TrackNumber) (model.Parcel, error)
	GetUserInfo(model.TrackNumber, model.TelegramID) (model.Parcel, bool, error)
	AddCheckpoint(model.TrackNumber, model.Checkpoint) error
	GetCheckpoints(trackNum model.TrackNumber, page uint64, pageSize uint64) ([]*model.Checkpoint, error)
	AddSubscription(model.TrackNumber, model.TelegramID) error
	DeleteSubscription(model.TrackNumber, model.TelegramID) error
	Exists(model.TrackNumber) (bool, error)
	Close()
}
