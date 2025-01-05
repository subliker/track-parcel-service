package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
)

func (s *store) AddSubscription(trackNum model.TrackNumber, userTID model.TelegramID) error {
	logger := s.logger.WithFields("command", "add subscription")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Insert("subscriptions").
		Columns("user_id", "parcel_track_number").
		Values(userTID, &trackNum).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of subscription inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	if _, err = s.db.Exec(query, args...); err != nil {
		if sqlErr, ok := err.(*pq.Error); ok {
			// if foreign key error
			if sqlErr.Code == "23053" {
				return parcel.ErrIncorrectForeignTrackNumber
			}
		}
		errMsg := fmt.Errorf("error executing of subscription inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	return nil
}

func (s *store) DeleteSubscription(trackNum model.TrackNumber, userTID model.TelegramID) error {
	logger := s.logger.WithFields("command", "delete subscription")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Delete("subscriptions").
		Where(squirrel.Eq{"parcel_track_number": trackNum, "user_id": userTID}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of subscription deleting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	res, err := s.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of subscription deleting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// check if was deleted
	aff, err := res.RowsAffected()
	if err != nil {
		errMsg := fmt.Errorf("error getting rows affected after subscription deleting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}
	if aff == 0 {
		return parcel.ErrParcelNotFound
	}

	return nil
}

func (s *store) ParcelSubscribers(trackNum model.TrackNumber) ([]model.TelegramID, error) {
	logger := s.logger.WithFields("command", "getting parcels subscribers")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("user_id").
		From("subscriptions").
		Where(squirrel.Eq{"parcel_track_number": trackNum}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of getting parcels subscribers: %s", err)
		logger.Error(errMsg)
		return nil, errMsg
	}

	// executing query
	rows, err := s.db.Query(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of getting parcels subscribers: %s", err)
		logger.Error(errMsg)
		return nil, errMsg
	}

	// making subscribers array
	sbs := make([]model.TelegramID, 0)

	// rows appending
	for rows.Next() {
		var userTID model.TelegramID
		if err := rows.Scan(&userTID); err != nil {
			errMsg := fmt.Errorf("failed to scan row: %s", err)
			logger.Error(errMsg)
			return nil, errMsg
		}
		sbs = append(sbs, userTID)
	}

	// check rows error
	if err := rows.Err(); err != nil {
		errMsg := fmt.Errorf("row iteration error: %s", err)
		logger.Error(errMsg)
		return nil, errMsg
	}
	return sbs, nil
}

func (s *store) GetSubscribed(trackNum model.TrackNumber, userTID model.TelegramID) (bool, error) {
	logger := s.logger.WithFields("command", "get info for user")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("1").
		From("subscriptions").
		Where(squirrel.Eq{"parcel_track_number": trackNum, "user_id": userTID}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of user parcel getting: %s", err)
		logger.Error(errMsg)
		return false, errMsg
	}

	// executing query
	r := 0
	err = s.db.QueryRow(query, args...).Scan(&r)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		errMsg := fmt.Errorf("error executing of getting subscribed: %s", err)
		logger.Error(errMsg)
		return false, errMsg
	}

	return true, nil
}
