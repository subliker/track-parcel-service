package auth

import (
	"context"
	"errors"

	ssov1 "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	serverApi struct {
		ssov1.UnimplementedAuthServer
		auth Auth
	}

	Auth interface {
		RegisterTelegramID(ctx context.Context, telegramId string) (userId string, err error)
		IsManager(ctx context.Context, telegramID string) (isManager bool, err error)
	}
)

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &serverApi{auth: auth})
}

func (s *serverApi) RegisterTelegramID(ctx context.Context, in *ssov1.RegisterTelegramIDRequest) (*ssov1.RegisterTelegramIDResponse, error) {
	if in.TelegramId == "" {
		return nil, status.Error(codes.InvalidArgument, "telegram id is empty")
	}

	uid, err := s.auth.RegisterTelegramID(ctx, in.GetTelegramId())
	if err != nil {
		if errors.Is(err, "") {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterTelegramIDResponse{UserId: uid}, nil
}
