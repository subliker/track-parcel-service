package user

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Register(ctx context.Context, in *pb.RegisterRequest) error
	GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error)
	Auth(ctx context.Context, in *pb.AuthRequest) error
	Close() error
}

type client struct {
	api     pb.UserClient
	ccClose func() error

	logger logger.Logger
}

func New(ctx context.Context, logger logger.Logger, cfg Config) (Client, error) {
	logger.Infof("grpc user client starts on target %s", cfg.Target)

	// opening client connection
	cc, err := grpc.NewClient(cfg.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &client{}, err
	}

	c := client{
		api:     pb.NewUserClient(cc),
		logger:  logger.WithFields("layer", "grpc client", "service", "user", "target", cc.Target()),
		ccClose: cc.Close,
	}

	return &c, nil
}

func (c *client) Close() error {
	if err := c.ccClose(); err != nil {
		return fmt.Errorf("error closing grpc user client connection: %s", err)
	}
	return nil
}
