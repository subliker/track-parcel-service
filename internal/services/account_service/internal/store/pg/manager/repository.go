package manager

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
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

func (r *Repository) Register(manager model.Manager) error {
	logger := r.logger.WithFields("command", "register")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// generate api token
	apiToken, err := model.NewManagerApiToken()
	if err != nil {
		errMsg := fmt.Errorf("error making manager api token inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// build query
	query, args, err := psql.Insert("managers").
		Columns("telegram_id", "full_name", "phone_number", "email", "company", "api_token").
		Values(&manager.TelegramId, &manager.FullName, &manager.PhoneNumber, &manager.Email, &manager.Company, &apiToken).
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

func (r *Repository) Get(tID model.TelegramID) (model.Manager, error) {
	logger := r.logger.WithFields("command", "get")

	// making manager struct
	var manager model.Manager

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
		return manager, errMsg
	}

	// executing query
	row := r.db.QueryRow(query, args...)
	err = row.Scan(&manager.TelegramId, &manager.FullName, &manager.PhoneNumber, &manager.Email, &manager.Company)
	if errors.Is(err, sql.ErrNoRows) {
		return manager, store.ErrManagerNotFound
	} else if err != nil {
		return manager, err
	}

	return manager, nil
}

func (r *Repository) GetApiToken(tID model.TelegramID) (model.ManagerApiToken, error) {
	logger := r.logger.WithFields("command", "get api token")

	// making api token var
	var apiToken model.ManagerApiToken

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
		return apiToken, errMsg
	}

	// executing query
	row := r.db.QueryRow(query, args...)
	err = row.Scan(&apiToken)
	if errors.Is(err, sql.ErrNoRows) {
		return apiToken, store.ErrManagerNotFound
	} else if err != nil {
		return apiToken, err
	}

	return apiToken, nil
}
