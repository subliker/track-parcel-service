package user

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/models/user"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerApi struct {
	pb.UnimplementedUserServer

	store store.Store

	logger logger.Logger
}

// New creates new instance of user server api
func New(logger logger.Logger, store store.Store) *ServerApi {
	logger = logger.WithFields("layer", "user server api")
	return &ServerApi{
		store:  store,
		logger: logger,
	}
}

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "register")

	// add user to store
	if err := s.store.User().Register(user.User{
		TelegramId:  telegram.ID(req.UserTelegramId),
		FullName:    req.UserFullName,
		PhoneNumber: req.UserPhoneNumber,
	}); err != nil {
		err = fmt.Errorf("error register user(%d): %s", req.UserTelegramId, err)
		logger.Error(err)
		return nil, err
	}

	return nil, nil
}

func (s *ServerApi) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	logger := s.logger.WithFields("handler", "get info")

	// getting user from repo
	u, err := s.store.User().Get(telegram.ID(req.UserTelegramId))
	if err != nil {
		err = fmt.Errorf("error getting user(%d): %s", req.UserTelegramId, err)
		logger.Error(err)
		return nil, err
	}

	return &pb.GetInfoResponse{
		UserFullName:    u.FullName,
		UserPhoneNumber: u.PhoneNumber,
	}, nil
}
