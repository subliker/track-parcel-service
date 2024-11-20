package user

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "register")
	const errMsg = "error register user(%d): %s"

	// add user to store
	if err := s.repo.Register(model.User{
		TelegramId:  model.TelegramID(req.UserTelegramId),
		FullName:    req.UserFullName,
		PhoneNumber: req.UserPhoneNumber,
	}); err != nil {
		errMsg := fmt.Sprintf(errMsg, req.UserTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	logger := s.logger.WithFields("handler", "get info")
	const errMsg = "error getting user(%d): %s"

	// getting user from repo
	u, err := s.repo.Get(model.TelegramID(req.UserTelegramId))
	if err == store.ErrUserNotFound {
		errMsg := fmt.Sprintf(errMsg, req.UserTelegramId, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.UserTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.GetInfoResponse{
		UserFullName:    u.FullName,
		UserPhoneNumber: u.PhoneNumber,
	}, nil
}
