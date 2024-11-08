package session

import (
	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
)

type Store interface {
	Add(tID telegram.ID) error
	Remove(tID telegram.ID) error
	Get(tID telegram.ID) (Session, error)
	Contains(tID telegram.ID) bool
}

type Session interface {
	State() interface{}
	SetState(state interface{})
	ClearState()
}
