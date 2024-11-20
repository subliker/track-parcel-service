package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/models/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

const tableName = "users"

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

func (r *Repository) Register(user user.User) error {
	logger := r.logger.WithFields("command", "register")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Insert("users").
		Columns("telegram_id", "full_name", "phone_number").
		Values(user.TelegramId, user.FullName, user.PhoneNumber).
		ToSql()
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

func (r *Repository) Get(tID telegram.ID) (user.User, error) {
	logger := r.logger.WithFields("command", "get")

	// making user struct
	var u user.User

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("telegram_id", "full_name", "phone_number").
		From("users").
		Where(squirrel.Eq{"telegram_id": tID}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of user selecting: %s", err)
		logger.Error(errMsg)
		return u, errMsg
	}

	// executing query
	row := r.db.QueryRow(query, args...)
	err = row.Scan(&u.TelegramId, &u.FullName, &u.PhoneNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return u, store.ErrUserNotFound
	} else if err != nil {
		return u, err
	}

	return u, nil
}
