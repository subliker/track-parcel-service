package user

import (
	"database/sql"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

type Repository struct {
	db     *sql.DB
	logger logger.Logger
}

func New(logger logger.Logger, db *sql.DB) store.UserRepository {

	return &Repository{
		db:     db,
		logger: logger.WithFields("layer", "user repository"),
	}
}
