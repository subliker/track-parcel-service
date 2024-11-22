package session

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

type Store interface {
	Add(tID model.TelegramID) error
	Remove(tID model.TelegramID) error
	Get(tID model.TelegramID) (Session, error)
	Contains(tID model.TelegramID) bool
	Ensure(tID model.TelegramID) error
	EnsureGet(tID model.TelegramID) (Session, error)
}

type Session interface {
	State() interface{}
	SetState(state interface{})
	ClearState()
}
