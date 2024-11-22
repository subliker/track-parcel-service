package pm

import (
	"context"

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
}

type client struct {
	api    pb.ParcelsManagerClient
	logger logger.Logger
}

func New(ctx context.Context, logger logger.Logger, cfg Config) Client {
	logger.Infof("grpc parcels manager client starts on target %s", cfg.Target)
	cc, err := grpc.NewClient(cfg.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(err)
	}

	return &client{
		api:    pb.NewParcelsManagerClient(cc),
		logger: logger.WithFields("layer", "grpc client", "service", "manager", "target", cc.Target()),
	}
}
