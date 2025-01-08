package pg

import (
	"fmt"
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
)

func (s *store) AddCheckpoint(tNum model.TrackNumber, cp model.Checkpoint) error {
	logger := s.logger.WithFields("command", "add checkpoint")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Insert("checkpoints").
		Columns("time", "place", "description", "parcel_track_number", "parcel_status").
		Values(&cp.Time, &cp.Place, &cp.Description, &tNum, &cp.ParcelStatus).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of checkpoint inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	if _, err := s.db.Exec(query, args...); err != nil {
		if sqlErr, ok := err.(*pq.Error); ok {
			// if foreign key error
			if sqlErr.Code == "23053" {
				return parcel.ErrIncorrectForeignTrackNumber
			}
		}
		errMsg := fmt.Errorf("error executing of checkpoint inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	return nil
}

func (s *store) GetCheckpoints(trackNum model.TrackNumber, page uint64, pageSize uint64) ([]*model.Checkpoint, error) {
	logger := s.logger.WithFields("command", "get checkpoints")

	// making query builder
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// build query
	query, args, err := psql.Select("time", "place", "description", "parcel_status").
		From("checkpoints").
		Where(squirrel.Eq{"parcel_track_number": trackNum}).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		ToSql()
	if err != nil {
		errMsg := fmt.Errorf("error making query of getting checkpoints: %s", err)
		logger.Error(errMsg)
		return nil, errMsg
	}

	// executing query
	rows, err := s.db.Query(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of getting checkpoints: %s", err)
		logger.Error(errMsg)
		return nil, errMsg
	}

	// making checkpoints array
	cps := make([]*model.Checkpoint, 0)

	// rows appending
	for rows.Next() {
		var cp model.Checkpoint
		if err := rows.Scan(&cp.Time, &cp.Place, &cp.Description); err != nil {
			errMsg := fmt.Errorf("failed to scan row: %s", err)
			logger.Error(errMsg)
			return nil, errMsg
		}
		cps = append(cps, &cp)
	}

	// check rows error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %s", err)
	}

	return cps, nil
}
