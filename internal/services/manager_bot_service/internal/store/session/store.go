package session

import (
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	models "github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
)

var sessionsPool = sync.Pool{
	New: func() any { return &userSession{} },
}

type store struct {
	cache *expirable.LRU[models.TelegramID, *userSession]
}

func New() session.Store {
	var s store

	// creating new cache
	s.cache = expirable.NewLRU(1024, s.handleAutoClear, time.Hour*48)

	return &s
}

func (s *store) handleAutoClear(tID models.TelegramID, ss *userSession) {
	logger.Zap.Infof("%s was wiped", tID)
}

func (s *store) Contains(tID models.TelegramID) bool {
	return s.cache.Contains(tID)
}

func (s *store) Add(tID models.TelegramID) error {
	// check if session is already exist
	ok := s.cache.Contains(tID)
	if ok {
		return session.ErrSessionIsAlreadyExist
	}

	// getting new instance of session from sync pool
	ns := sessionsPool.Get().(*userSession)

	// saving session in cache
	s.cache.Add(tID, ns)
	return nil
}

func (s *store) Remove(tID models.TelegramID) error {
	// check if session is already exist
	ss, ok := s.cache.Get(tID)
	if !ok {
		return session.ErrSessionIsNotExist
	}

	// return instance of session to sync pool
	sessionsPool.Put(ss)

	// removing session from cache
	s.cache.Remove(tID)
	return nil
}

func (s *store) Get(tID models.TelegramID) (session.Session, error) {
	// getting session
	ss, ok := s.cache.Get(tID)
	if !ok {
		return nil, session.ErrSessionIsNotExist
	}
	return ss, nil
}
