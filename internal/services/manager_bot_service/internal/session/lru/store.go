package lru

import (
	"errors"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
)

var sessionsPool = sync.Pool{
	New: func() any { return &userSession{} },
}

type store struct {
	cache  *expirable.LRU[model.TelegramID, *userSession]
	logger logger.Logger
}

// New creates new session store
func New(logger logger.Logger) session.Store {
	var s store

	// creating new cache
	s.cache = expirable.NewLRU(1024, s.handleAutoClear, time.Hour*48)

	// set logger
	s.logger = logger.WithFields("layer", "session store")

	return &s
}

func (s *store) handleAutoClear(tID model.TelegramID, ss *userSession) {
	s.logger.Infof("%s was wiped", tID)
}

func (s *store) Contains(tID model.TelegramID) bool {
	return s.cache.Contains(tID)
}

func (s *store) Add(tID model.TelegramID) error {
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

func (s *store) Remove(tID model.TelegramID) error {
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

func (s *store) Get(tID model.TelegramID) (session.Session, error) {
	// getting session
	ss, ok := s.cache.Get(tID)
	if !ok {
		return nil, session.ErrSessionIsNotExist
	}
	return ss, nil
}

func (s *store) Ensure(tID model.TelegramID) error {
	// try to add session
	err := s.Add(tID)
	if !errors.Is(err, session.ErrSessionIsAlreadyExist) {
		return err
	}

	return nil
}

func (s *store) EnsureGet(tID model.TelegramID) (session.Session, error) {
	// try to add session
	err := s.Add(tID)
	if !errors.Is(err, session.ErrSessionIsAlreadyExist) {
		return nil, err
	}
}
