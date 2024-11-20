package pg

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store/pg/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store/pg/user"
)

var migrateMode bool

func init() {
	flag.BoolVar(&migrateMode, "migrate", false, "set to use migrations")
}

type pgStore struct {
	db      *sql.DB
	user    store.UserRepository
	manager store.ManagerRepository
}

func New(logger logger.Logger, cfg config.DBConfig) (store.Store, error) {
	logger = logger.WithFields("layer", "pgstore")

	// pgs connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

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

	// migrations
	if migrateMode {
		if err := goose.Up(db, "migrations"); err != nil {
			logger.Fatalf("migrations error: %s", err)
		}
		logger.Info("migration was successful")
	}

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
