package manager

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerApi struct {
	pb.UnimplementedManagerServer

	repo   store.ManagerRepository
	logger logger.Logger
}

// New creates new instance of manager server api
func New(logger logger.Logger, repo store.ManagerRepository) *ServerApi {
	return &ServerApi{
		repo:   repo,
		logger: logger.WithFields("layer", "manager server api"),
	}
}

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "register")
	const errMsg = "error register manager(%d): %s"

	// add manager to store
	if err := s.repo.Register(model.Manager{
		TelegramId:  model.TelegramID(req.ManagerTelegramId),
		FullName:    req.ManagerFullName,
		Email:       req.ManagerEmail,
		PhoneNumber: req.ManagerPhoneNumber,
		Company:     req.ManagerCompany,
	}); err != nil {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	logger := s.logger.WithFields("handler", "get info")
	const errMsg = "error getting manager(%d): %s"

	// getting manager from repo
	m, err := s.repo.Get(model.TelegramID(req.ManagerTelegramId))
	if err == store.ErrManagerNotFound {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.GetInfoResponse{
		ManagerFullName:    m.FullName,
		ManagerEmail:       m.Email,
		ManagerPhoneNumber: m.PhoneNumber,
		ManagerCompany:     m.Company,
	}, nil
}

func (s *ServerApi) GetApiToken(ctx context.Context, req *pb.GetApiTokenRequest) (*pb.GetApiTokenResponse, error) {
	logger := s.logger.WithFields("handler", "get api token")

	// getting api token from repo
	const errMsg = "error getting manager(%d) api token: %s"
	t, err := s.repo.GetApiToken(model.TelegramID(req.ManagerTelegramId))
	if err == store.ErrManagerNotFound {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.GetApiTokenResponse{
		ManagerApiToken: fmt.Sprint(t),
	}, nil
}
