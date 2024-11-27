package pu

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pupb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	GetParcel(ctx context.Context, in *pb.GetParcelRequest) (*pb.GetParcelResponse, error)
	GetCheckpoints(ctx context.Context, in *pb.GetCheckpointsRequest) (*pb.GetCheckpointsResponse, error)
	Close() error
}

type client struct {
	api     pb.ParcelsUserClient
	ccClose func() error

	logger logger.Logger
}

func New(ctx context.Context, logger logger.Logger, cfg Config) (Client, error) {
	logger.Infof("grpc parcels manager client starts on target %s", cfg.Target)
	// opening client connection
	cc, err := grpc.NewClient(cfg.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(err)
	}

	return &client{
		api:     pb.NewParcelsUserClient(cc),
		logger:  logger.WithFields("layer", "grpc client", "service", "parcels user", "target", cc.Target()),
		ccClose: cc.Close,
	}, nil
}

func (c *client) Close() error {
	if err := c.ccClose(); err != nil {
		return fmt.Errorf("error closing grpc user client connection: %s", err)
	}
	return nil
}
