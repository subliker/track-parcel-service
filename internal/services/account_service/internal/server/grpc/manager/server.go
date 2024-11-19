package manager

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/models/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerApi struct {
	pb.UnimplementedManagerServer

	store  store.Store
	logger logger.Logger
}

// New creates new instance of manager server api
func New(logger logger.Logger, store store.Store) *ServerApi {
	logger = logger.WithFields("layer", "manager server api")
	return &ServerApi{
		store:  store,
		logger: logger,
	}
}

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "register")

	// add manager to store
	if err := s.store.Manager().Register(manager.Manager{
		TelegramId:  telegram.ID(req.ManagerTelegramId),
		FullName:    req.ManagerFullName,
		Email:       req.ManagerEmail,
		PhoneNumber: req.ManagerPhoneNumber,
		Company:     req.ManagerCompany,
	}); err != nil {
		err = fmt.Errorf("error register manager(%d): %s", req.ManagerTelegramId, err)
		logger.Error(err)
		return nil, err
	}

	return nil, nil
}

func (s *ServerApi) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	logger := s.logger.WithFields("handler", "get info")

	// getting manager from repo
	m, err := s.store.Manager().Get(telegram.ID(req.ManagerTelegramId))
	if err != nil {
		err = fmt.Errorf("error getting manager(%d): %s", req.ManagerTelegramId, err)
		logger.Error(err)
		return nil, err
	}

	return &pb.GetInfoResponse{
		ManagerFullName:    m.FullName,
		ManagerEmail:       m.Email,
		ManagerPhoneNumber: m.PhoneNumber,
		ManagerCompany:     m.Company,
	}, nil
}
