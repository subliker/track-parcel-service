package parcel

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

type Store interface {
	Add(model.Parcel) (model.TrackNumber, error)
	Delete(model.TrackNumber) error
	GetInfo(model.TrackNumber) (model.Parcel, error)
	Close()
}
