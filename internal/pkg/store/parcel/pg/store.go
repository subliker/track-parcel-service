package pg

import (
	"database/sql"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
)

type store struct {
	db     *sql.DB
	logger logger.Logger
}

// New creates new instance of postgres parcel store
func New(logger logger.Logger, cfg Config) (parcel.Store, error) {
	logger = logger.WithFields("layer", "parcel postgres store")

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

	return &store{
		db:     db,
		logger: logger,
	}, nil
}

func (s *store) Close() error {
	defer s.logger.Info("store stopped")
	return s.db.Close()
}
