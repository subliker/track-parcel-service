package pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
)

func (s *store) CheckAccess(trackNum model.TrackNumber, tID model.TelegramID) (bool, error) {
	p, err := s.GetInfo(trackNum)
	if err != nil {
		return false, err
	}
	return p.ManagerID == tID, nil
}

func (s *store) Add(p model.Parcel) (model.TrackNumber, error) {
	logger := s.logger.WithFields("command", "add")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// generating track number
	trackNum := model.NewTrackNumber()

	// build query
	query, args, err := psql.Insert("parcels").
		Columns("track_number", "name", "manager_id", "recipient", "arrival_address", "forecast_date", "description", "status").
		Values(&trackNum, &p.Name, &p.ManagerID, &p.Recipient, &p.ArrivalAddress, &p.ForecastDate, &p.Description, &p.Status).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of parcel inserting: %s", err)
		logger.Error(errMsg)
		return "", errMsg
	}

	// executing query
	res, err := s.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of parcel inserting: %s", err)
		logger.Error(errMsg)
		return "", errMsg
	}

	logger.Info(res)
	return trackNum, nil
}

func (s *store) Delete(trackNum model.TrackNumber) error {
	logger := s.logger.WithFields("command", "delete")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Delete("parcels").
		Where(squirrel.Eq{"track_number": trackNum}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of parcel deleting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	res, err := s.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of parcel deleting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// check if was deleted
	aff, err := res.RowsAffected()
	if err != nil {
		errMsg := fmt.Errorf("error getting rows affected after parcel deleting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}
	if aff == 0 {
		return parcel.ErrParcelNotFound
	}

	return nil
}

func (s *store) GetInfo(trackNum model.TrackNumber) (model.Parcel, error) {
	logger := s.logger.WithFields("command", "get info")

	// making parcel struct
	var p model.Parcel

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("name", "manager_id", "recipient", "arrival_address", "forecast_date", "description", "status").
		From("parcels").
		Where(squirrel.Eq{"track_number": trackNum}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of parcel getting: %s", err)
		logger.Error(errMsg)
		return p, errMsg
	}

	// executing query
	row := s.db.QueryRow(query, args...)
	err = row.Scan(&p.Name, &p.ManagerID, &p.Recipient, &p.ArrivalAddress, &p.ForecastDate, &p.Description, &p.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return p, parcel.ErrParcelNotFound
	}
	if err != nil {
		return p, err
	}

	return p, nil
}

func (s *store) GetUserInfo(trackNum model.TrackNumber, userTID model.TelegramID) (model.Parcel, bool, error) {
	logger := s.logger.WithFields("command", "get user info")

	// making parcel struct
	var p model.Parcel
	subscribed := false

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.
		Select(
			"parcels.name",
			"parcels.recipient",
			"parcels.arrival_address",
			"parcels.forecast_date",
			"parcels.description",
			"parcels.status",
			"subscriptions.id AS subscription_id",
		).
		From("parcels").
		LeftJoin("subscriptions ON parcels.track_number = subscriptions.parcel_track_number AND subscriptions.user_id = ?", userTID).
		Where(squirrel.Eq{"parcels.track_number": trackNum}).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of user parcel getting: %s", err)
		logger.Error(errMsg)
		return p, false, errMsg
	}

	// executing query
	row := s.db.QueryRow(query, args...)
	err = row.Scan(&p.Name, &p.Recipient, &p.ArrivalAddress, &p.ForecastDate, &p.Description, &p.Status, &subscribed)
	if errors.Is(err, sql.ErrNoRows) {
		return p, false, parcel.ErrParcelNotFound
	}
	if err != nil {
		return p, false, err
	}

	return p, subscribed, nil
}

func (s *store) Exists(trackNum model.TrackNumber) (bool, error) {
	logger := s.logger.WithFields("command", "exists")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.
		Select("1").
		Prefix("SELECT EXISTS (").
		From("parcels").
		Where(squirrel.Eq{"track_number": trackNum}).
		Suffix(")").
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of parcel exists selecting: %s", err)
		logger.Error(errMsg)
		return false, errMsg
	}

	// executing query
	var exists bool
	err = s.db.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		errMsg := fmt.Errorf("error execution query of parcel exists selecting: %s", err)
		logger.Error(errMsg)
		return false, errMsg
	}

	return exists, nil
}
