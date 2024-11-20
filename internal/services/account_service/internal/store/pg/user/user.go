package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

func (r *Repository) Register(user model.User) error {
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

func (r *Repository) Get(tID model.TelegramID) (model.User, error) {
	logger := r.logger.WithFields("command", "get")

	// making user struct
	var u model.User

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
	}
	if err != nil {
		return u, err
	}

	return u, nil
}