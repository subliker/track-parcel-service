package pu

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pupb"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerApi) GetParcel(ctx context.Context, req *pb.GetParcelRequest) (*pb.GetParcelResponse, error) {
	logger := s.logger.WithFields("handler", "get parcel")
	const errMsg = "error get parcel(%s): %s"

	// get parcel from store
	p, subscribed, err := s.store.GetUserInfo(model.TrackNumber(req.TrackNumber), model.TelegramID(req.UserTelegramId))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.GetParcelResponse{
		ParcelName:           p.Name,
		ParcelRecipient:      p.Recipient,
		ParcelArrivalAddress: p.ArrivalAddress,
		ParcelForecastDate:   timestamppb.New(p.ForecastDate),
		ParcelDescription:    p.Description,
		ParcelStatus:         pb.ParcelStatus(pb.ParcelStatus_value[string(p.Status)]),
		UserSubscribed:       subscribed,
	}, nil
}
