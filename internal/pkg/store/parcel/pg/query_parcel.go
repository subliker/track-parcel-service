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
		Columns("track_number", "name", "manager_id", "recipient", "arrival_address", "forecast_date", "description").
		Values(&trackNum, &p.Name, &p.ManagerID, &p.Recipient, &p.ArrivalAddress, &p.ForecastDate, &p.Description).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of parcel inserting: %s", err)
		logger.Error(errMsg)
		return "", errMsg
	}

	// executing query
	_, err = s.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of parcel inserting: %s", err)
		logger.Error(errMsg)
		return "", errMsg
	}

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

	// build subquery for last status
	subQuery, subArgs, err := squirrel.Select(
		"c.parcel_track_number",
		"c.parcel_status",
		"ROW_NUMBER() OVER (PARTITION BY c.parcel_track_number ORDER BY c.time DESC) AS rnk",
	).From("checkpoints c").
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making subquery of parcel getting: %s", err)
		logger.Error(errMsg)
		return p, errMsg
	}

	// build query
	query, args, err := psql.Select("name", "manager_id", "recipient", "arrival_address", "forecast_date", "description", "c.parcel_status as status").
		From("parcels p").
		Where(squirrel.Eq{"track_number": trackNum}).
		LeftJoin(fmt.Sprintf("(%s) AS c ON p.track_number = c.parcel_track_number AND c.rnk = 1", subQuery)).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of parcel getting: %s", err)
		logger.Error(errMsg)
		return p, errMsg
	}

	args = append(args, subArgs...)

	// executing query
	row := s.db.QueryRow(query, args...)
	err = row.Scan(&p.Name, &p.ManagerID, &p.Recipient, &p.ArrivalAddress, &p.ForecastDate, &p.Description, &p.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return p, parcel.ErrParcelNotFound
	}
	if err != nil {
		errMsg := fmt.Errorf("error executing of getting parcel: %s", err)
		logger.Error(errMsg)
		return p, errMsg
	}

	return p, nil
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
