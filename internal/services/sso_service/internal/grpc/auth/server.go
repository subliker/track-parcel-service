package auth

import (
	"context"
	"fmt"

	ssov1 "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/sso"
	"google.golang.org/grpc"
)

type (
	ServerApi struct {
		ssov1.UnimplementedAuthServer
		auth Auth
	}

	Auth interface {
		RegisterTelegramID(ctx context.Context, telegramId string) (userId string, err error)
		IsManager(ctx context.Context, telegramID string) (isManager bool, err error)
	}
)

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &ServerApi{auth: auth})
}

func (s *ServerApi) RegisterTelegramID(ctx context.Context, req *ssov1.RegisterTelegramIDRequest) (*ssov1.RegisterTelegramIDResponse, error) {
	// if req.TelegramId == "" {
	// 	return nil, status.Error(codes.InvalidArgument, "telegram id is empty")
	// }

	// uid, err := s.auth.RegisterTelegramID(ctx, req.GetTelegramId())
	// if err != nil {
	// 	if errors.Is(err, "") {
	// 		return nil, status.Error(codes.AlreadyExists, "user already exists")
	// 	}

	// 	return nil, status.Error(codes.Internal, "failed to register user")
	// }
	fmt.Print("it works")

	return &ssov1.RegisterTelegramIDResponse{UserId: ""}, nil
}
