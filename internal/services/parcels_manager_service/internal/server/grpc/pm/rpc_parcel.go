package pm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/gen/parcelpb"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/pmpb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerApi) AddParcel(ctx context.Context, req *pmpb.AddParcelRequest) (*pmpb.AddParcelResponse, error) {
	logger := s.logger.WithFields("handler", "add parcel")
	const errMsg = "error add parcel: %s"

	// add parcel to store
	trackNum, err := s.store.Add(model.Parcel{
		Name:           req.Parcel.Name,
		ManagerID:      model.TelegramID(req.Parcel.ManagerTelegramId),
		Recipient:      req.Parcel.Recipient,
		ArrivalAddress: req.Parcel.ArrivalAddress,
		ForecastDate:   req.Parcel.ForecastDate.AsTime(),
		Description:    req.Parcel.Description,
		Status:         model.Status(req.Parcel.Status.String()),
	})
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, err)
		logger.Error(errMsg)
		return &pmpb.AddParcelResponse{}, status.Error(codes.Internal, errMsg)
	}

	// add empty status
	err = s.store.AddCheckpoint(trackNum, model.Checkpoint{
		Time:         time.Now(),
		Place:        "",
		Description:  "Parcel was created",
		ParcelStatus: model.StatusPending,
	})
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, err)
		logger.Error(errMsg)
		return &pmpb.AddParcelResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &pmpb.AddParcelResponse{
		TrackNumber: string(trackNum),
	}, nil
}

func (s *ServerApi) DeleteParcel(ctx context.Context, req *pmpb.DeleteParcelRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "delete parcel")
	const errMsg = "error delete parcel(%s): %s"

	// delete parcel from store
	err := s.store.Delete(model.TrackNumber(req.TrackNumber))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) GetParcel(ctx context.Context, req *pmpb.GetParcelRequest) (*pmpb.GetParcelResponse, error) {
	logger := s.logger.WithFields("handler", "get parcel info")
	const errMsg = "error get parcel(%s): %s"

	// get parcel from store
	p, err := s.store.GetInfo(model.TrackNumber(req.TrackNumber))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pmpb.GetParcelResponse{
		Parcel: &parcelpb.Parcel{
			Name:              p.Name,
			ManagerTelegramId: int64(p.ManagerID),
			Recipient:         p.Recipient,
			ArrivalAddress:    p.ArrivalAddress,
			ForecastDate:      timestamppb.New(p.ForecastDate),
			Description:       p.Description,
			Status:            parcelpb.Status(parcelpb.Status_value[string(p.Status)]),
		},
	}, nil
}
