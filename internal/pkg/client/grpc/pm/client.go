package pm

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pmpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	AddParcel(ctx context.Context, in *pb.AddParcelRequest) (*pb.AddParcelResponse, error)
	DeleteParcel(ctx context.Context, in *pb.DeleteParcelRequest) error
	GetParcel(ctx context.Context, in *pb.GetParcelRequest) (*pb.GetParcelResponse, error)
	AddCheckpoint(ctx context.Context, in *pb.AddCheckpointRequest) error
	GetCheckpoints(ctx context.Context, in *pb.GetCheckpointsRequest) (*pb.GetCheckpointsResponse, error)
	Close() error
}

type client struct {
	api     pb.ParcelsManagerClient
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

	// // try to reach grpc server
	// if cc.GetState() != connectivity.Ready {
	// 	return nil, errors.New("grpc parcels manager server is not reachable")
	// }

	return &client{
		api:     pb.NewParcelsManagerClient(cc),
		logger:  logger.WithFields("layer", "grpc client", "service", "parcels manager", "target", cc.Target()),
		ccClose: cc.Close,
	}, nil
}

func (c *client) Close() error {
	if err := c.ccClose(); err != nil {
		return fmt.Errorf("error closing grpc manager client connection: %s", err)
	}
	return nil
}
