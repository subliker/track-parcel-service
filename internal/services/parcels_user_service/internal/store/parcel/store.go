package parcel

import "github.com/subliker/track-parcel-service/internal/pkg/model"

type UserStore interface {
	GetInfo(model.TrackNumber) (model.Parcel, error)
	GetCheckpoints(trackNum model.TrackNumber, page uint64, pageSize uint64) ([]*model.Checkpoint, error)
	Exists(model.TrackNumber) (bool, error)
	Close()
}
