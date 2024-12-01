package pg

import (
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
		Where(squirrel.Eq{"track_number": trackNum, "user_id": userTID}).
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
