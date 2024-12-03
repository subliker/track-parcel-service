package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store/pg/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store/pg/user"
)

type pgStore struct {
	db      *sql.DB
	user    store.UserRepository
	manager store.ManagerRepository
}

func New(logger logger.Logger, cfg Config) (store.Store, error) {
	logger = logger.WithFields("layer", "pgstore")

	// pgs connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)

	// opening sql connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// try to ping db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logger.Info("pgstore was connected")

	return &pgStore{
		db:      db,
		user:    user.New(logger, db),
		manager: manager.New(logger, db),
	}, nil
}

func (s *pgStore) Close() error {
	return s.db.Close()
}

func (s *pgStore) Manager() store.ManagerRepository {
	return s.manager
}

func (s *pgStore) User() store.UserRepository {
	return s.user
}
