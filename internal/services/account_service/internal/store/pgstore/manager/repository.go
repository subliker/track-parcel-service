package manager

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/models/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

type Repository struct {
	db     *sql.DB
	logger logger.Logger
}

func New(logger logger.Logger, db *sql.DB) store.ManagerRepository {

	return &Repository{
		db:     db,
		logger: logger.WithFields("layer", "manager repository"),
	}
}

func (r *Repository) Register(manager manager.Manager) error {
	logger := r.logger.WithFields("command", "register")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Insert("managers").
		Columns("telegram_id", "full_name", "phone_number", "email", "company").
		Values(&manager.TelegramId, &manager.FullName, &manager.PhoneNumber, &manager.Email, &manager.Company).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of manager inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	res, err := r.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of manager inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	logger.Info(res)
	return nil
}

func (r *Repository) Get(tID telegram.ID) (manager.Manager, error) {
	logger := r.logger.WithFields("command", "get")

	// making manager struct
	var m manager.Manager

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("telegram_id", "full_name", "phone_number", "email", "company").
		From("managers").
		Where(squirrel.Eq{"telegram_id": tID}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of manager selecting: %s", err)
		logger.Error(errMsg)
		return m, errMsg
	}

	// executing query
	row := r.db.QueryRow(query, args...)
	err = row.Scan(&m.TelegramId, &m.FullName, &m.PhoneNumber, &m.Email, &m.Company)
	if errors.Is(err, sql.ErrNoRows) {
		return m, ErrManagerNotFound
	} else if err != nil {
		return m, err
	}

	return m, nil
}

func (r *Repository) GetApiToken(tID telegram.ID) (manager.ApiToken, error) {
	logger := r.logger.WithFields("command", "get api token")

	// making api token var
	var t manager.ApiToken

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("api_token").
		From("managers").
		Where(squirrel.Eq{"telegram_id": tID}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of manager api token selecting: %s", err)
		logger.Error(errMsg)
		return t, errMsg
	}

	// executing query
	row := r.db.QueryRow(query, args...)
	err = row.Scan(&t)
	if errors.Is(err, sql.ErrNoRows) {
		return t, ErrManagerNotFound
	} else if err != nil {
		return t, err
	}

	return t, nil
}
