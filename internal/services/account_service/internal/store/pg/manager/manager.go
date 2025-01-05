package manager

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

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
		Values(&manager.TelegramID, &manager.FullName, &manager.PhoneNumber, &manager.Email, &manager.Company, &apiToken).
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
	err = row.Scan(&manager.TelegramID, &manager.FullName, &manager.PhoneNumber, &manager.Email, &manager.Company)
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

func (r *Repository) RetrieveManagerIdByApiKey(apiToken model.ManagerApiToken) (model.TelegramID, error) {
	logger := r.logger.WithFields("command", "get telegram id")

	// making telegram id
	var tID model.TelegramID

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("telegram_id").
		From("managers").
		Where(squirrel.Eq{"api_token": apiToken}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of get telegram id  selecting: %s", err)
		logger.Error(errMsg)
		return tID, errMsg
	}

	// executing query
	row := r.db.QueryRow(query, args...)
	err = row.Scan(&tID)
	if errors.Is(err, sql.ErrNoRows) {
		return tID, store.ErrManagerNotFound
	} else if err != nil {
		return tID, err
	}

	return tID, nil
}

func (r *Repository) Exists(tID model.TelegramID) (bool, error) {
	logger := r.logger.WithFields("command", "exists")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.
		Select("1").
		Prefix("SELECT EXISTS (").
		From("managers").
		Where(squirrel.Eq{"telegram_id": tID}).
		Suffix(")").
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of manager exists selecting: %s", err)
		logger.Error(errMsg)
		return false, errMsg
	}

	// executing query
	var exists bool
	err = r.db.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		errMsg := fmt.Errorf("error execution query of manager exists selecting: %s", err)
		logger.Error(errMsg)
		return false, errMsg
	}

	return exists, nil
}
