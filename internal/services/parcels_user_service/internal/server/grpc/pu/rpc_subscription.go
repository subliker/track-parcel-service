package pu

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/store/parcel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerApi) AddSubscription(ctx context.Context, req *pb.AddSubscriptionRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "add subscription")
	const errMsg = "error add subscription(%s): %s"

	// check if subscribed
	subscribed, err := s.store.GetSubscribed(model.TrackNumber(req.TrackNumber), model.TelegramID(req.UserTelegramId))
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}
	if subscribed {
		return nil, status.Error(codes.AlreadyExists, "")
	}

	// add subscription in store
	err = s.store.AddSubscription(model.TrackNumber(req.TrackNumber), model.TelegramID(req.UserTelegramId))
	if errors.Is(err, parcel.ErrIncorrectForeignTrackNumber) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil

}
func (s *ServerApi) DeleteSubscription(ctx context.Context, req *pb.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "delete subscription")
	const errMsg = "error delete subscription(%s): %s"

	// delete subscription from store
	err := s.store.DeleteSubscription(model.TrackNumber(req.TrackNumber), model.TelegramID(req.UserTelegramId))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil
}
