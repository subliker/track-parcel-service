package session

import models "github.com/subliker/track-parcel-service/internal/pkg/models/telegram"

type Store interface {
	Add(tID models.TelegramID) error
	Remove(tID models.TelegramID) error
	Get(tID models.TelegramID) (Session, error)
	Contains(tID models.TelegramID) bool
}

type Session interface {
}
