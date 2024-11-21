package manager

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "register")
	const errMsg = "error register manager(%d): %s"

	// check context
	select {
	case <-ctx.Done():
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, ctx.Err())
		logger.Info(errMsg)
		return nil, status.Error(codes.Canceled, errMsg)
	default:
	}

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

	// check context
	select {
	case <-ctx.Done():
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, ctx.Err())
		logger.Info(errMsg)
		return nil, status.Error(codes.Canceled, errMsg)
	default:
	}

	// getting manager from repo
	m, err := s.repo.Get(model.TelegramID(req.ManagerTelegramId))
	if err == store.ErrManagerNotFound {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
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
	const errMsg = "error getting manager(%d) api token: %s"

	// check context
	select {
	case <-ctx.Done():
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, ctx.Err())
		logger.Info(errMsg)
		return nil, status.Error(codes.Canceled, errMsg)
	default:
	}

	// getting api token from repo
	t, err := s.repo.GetApiToken(model.TelegramID(req.ManagerTelegramId))
	if err == store.ErrManagerNotFound {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
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

func (s *ServerApi) Auth(ctx context.Context, req *pb.AuthRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "auth")
	const errMsg = "error auth manager(%d): %s"

	// check context
	select {
	case <-ctx.Done():
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, ctx.Err())
		logger.Info(errMsg)
		return nil, status.Error(codes.Canceled, errMsg)
	default:
	}

	// check if manager exists in repo
	exists, err := s.repo.Exists(model.TelegramID(req.ManagerTelegramId))
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}
	if !exists {
		errMsg := fmt.Sprintf(errMsg, req.ManagerTelegramId, "manager wasn't found")
		return nil, status.Error(codes.NotFound, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) AuthApiToken(ctx context.Context, req *pb.AuthApiTokenRequest) (*pb.AuthApiTokenResponse, error) {
	logger := s.logger.WithFields("handler", "auth api token")
	const errMsg = "error auth api token: %s"

	// check context
	select {
	case <-ctx.Done():
		errMsg := fmt.Sprintf(errMsg, ctx.Err())
		logger.Info(errMsg)
		return nil, status.Error(codes.Canceled, errMsg)
	default:
	}

	// get manager telegram id
	tID, err := s.repo.GetTelegramId(model.ManagerApiToken(req.ManagerApiToken))
	if err == store.ErrManagerNotFound {
		errMsg := fmt.Sprintf(errMsg, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.AuthApiTokenResponse{
		ManagerTelegramId: int64(tID),
	}, nil
}
