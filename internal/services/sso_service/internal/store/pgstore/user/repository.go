package user

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/models/user"
)

const tableName = "users"

type Repository struct {
	db     *sql.DB
	logger logger.Logger
}

func New(db *sql.DB, logger logger.Logger) Repository {

	return Repository{
		db:     db,
		logger: logger.WithFields("layer", "user repository"),
	}
}

func (r *Repository) Register(user user.User) error {
	logger := r.logger.WithFields("command", "register")

	// making query
	query, args, err := squirrel.Insert(tableName).Columns("telegram_id").
		Values(user.TelegramId).ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of user inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	res, err := r.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of user inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	logger.Info(res)
	return nil
}
