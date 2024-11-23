package manager

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Register(ctx context.Context, in *pb.RegisterRequest) error
	GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error)
	GetApiToken(ctx context.Context, in *pb.GetApiTokenRequest) (*pb.GetApiTokenResponse, error)
	Auth(ctx context.Context, in *pb.AuthRequest) error
	AuthApiToken(ctx context.Context, in *pb.AuthApiTokenRequest) (*pb.AuthApiTokenResponse, error)
}

type client struct {
	api    pb.ManagerClient
	logger logger.Logger
}

func New(ctx context.Context, logger logger.Logger, cfg Config) Client {
	logger.Infof("grpc manager client starts on target %s", cfg.Target)
	cc, err := grpc.NewClient(cfg.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(err)
	}

	return &client{
		api:    pb.NewManagerClient(cc),
		logger: logger.WithFields("layer", "grpc client", "service", "manager", "target", cc.Target()),
	}
}