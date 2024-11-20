package pg

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/pressly/goose/v3"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
)

var migrateMode bool

func init() {
	flag.BoolVar(&migrateMode, "migrate", false, "set to use migrations")
}

type store struct {
	db     *sql.DB
	logger logger.Logger
}

// New creates new instance of postgres parcel store
func New(logger logger.Logger, cfg config.DBConfig) (parcel.Store, error) {
	logger = logger.WithFields("layer", "parcel postgres store")

	// pgs connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// opening sql connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// try to ping db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logger.Info("pgstore was connected")

	// migrations
	if migrateMode {
		if err := goose.Up(db, "migrations"); err != nil {
			logger.Fatalf("migrations error: %s", err)
		}
		logger.Info("migration was successful")
	}

	return &store{
		db:     db,
		logger: logger,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
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
		errMsg := fmt.Errorf("error making query of parcel inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// executing query
	res, err := s.db.Exec(query, args...)
	if err != nil {
		errMsg := fmt.Errorf("error executing of parcel inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}

	// check if was deleted
	aff, err := res.RowsAffected()
	if err != nil {
		errMsg := fmt.Errorf("error getting rows affected after parcel inserting: %s", err)
		logger.Error(errMsg)
		return errMsg
	}
	if aff == 0 {
		return parcel.ErrParcelNotFound
	}

	logger.Info(res)
	return nil
}

func (s *store) GetInfo(trackNum model.TrackNumber) (model.Parcel, error) {
	logger := s.logger.WithFields("command", "get")

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
		errMsg := fmt.Errorf("error making query of parcel inserting: %s", err)
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
