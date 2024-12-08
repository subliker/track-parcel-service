package pu

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/gen/parcelpb"
	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerApi) GetParcel(ctx context.Context, req *pb.GetParcelRequest) (*pb.GetParcelResponse, error) {
	logger := s.logger.WithFields("handler", "get parcel")
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

	// getting subscribed
	subscribed, err := s.store.GetSubscribed(model.TrackNumber(req.TrackNumber), model.TelegramID(req.UserTelegramId))
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.GetParcelResponse{
		Parcel: &parcelpb.Parcel{
			Name:           p.Name,
			Recipient:      p.Recipient,
			ArrivalAddress: p.ArrivalAddress,
			ForecastDate:   timestamppb.New(p.ForecastDate),
			Description:    p.Description,
			Status:         parcelpb.Status(parcelpb.Status_value[string(p.Status)]),
		},
		UserSubscribed: subscribed,
	}, nil
}
