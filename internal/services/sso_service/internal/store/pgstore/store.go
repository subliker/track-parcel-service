package pgstore

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store/pgstore/manager"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store/pgstore/user"
)

var migrateMode bool

func init() {
	flag.BoolVar(&migrateMode, "migrate", false, "set to use migrations")
}

type pgStore struct {
	db      *sql.DB
	user    *user.Repository
	manager *manager.Repository
}

func New(cfg config.DBConfig) (store.Store, error) {
	logger := zap.Logger.WithFields("layer", "pgstore")

	// pgs connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// opening sql connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	logger.Info("pgstore was connected")

	// migrations
	if migrateMode {
		if err := goose.Up(db, "migrations"); err != nil {
			logger.Fatal("migrations error: %s", err)
		}
	}

	userRepo := user.New(db, logger)

	return &pgStore{
		db:   db,
		user: &userRepo,
	}, nil
}

func (s *pgStore) Close() error {
	return s.db.Close()
}

func (s *pgStore) Manager() (*manager.Repository, error) {
	if s.manager == nil {
		return nil, fmt.Errorf("manager repository is not set")
	}
	return s.manager, nil
}

func (s *pgStore) User() (*user.Repository, error) {
	if s.user == nil {
		return nil, fmt.Errorf("user repository is not set")
	}
	return s.user, nil
}
