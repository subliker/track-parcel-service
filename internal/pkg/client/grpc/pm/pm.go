package pm

import (
	"context"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pmpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *client) AddParcel(ctx context.Context, in *pb.AddParcelRequest) (*pb.AddParcelResponse, error) {
	logger := c.logger.WithFields("request", "add parcel")
	const errMsg = "request add parcel error: %s"

	// api call
	res, err := c.api.AddParcel(ctx, in)
	if err == nil {
		return res, nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.AlreadyExists:
			return ErrManagerIsAlreadyExist
		case codes.Internal:
			logger.Error(errMsg)
			return ErrInternal
		default:
			logger.Error(errMsg)
			return ErrUnexpected
		}
	}
	logger.Errorf(errMsg, err)
	return ErrUnexpected
}
func (c *client) DeleteParcel(ctx context.Context, in *pb.DeleteParcelRequest) error {

}
func (c *client) GetParcel(ctx context.Context, in *pb.GetParcelRequest) (*pb.GetParcelResponse, error) {

}
func (c *client) AddCheckpoint(ctx context.Context, in *pb.AddCheckpointRequest) error {

}
func (c *client) GetCheckpoints(ctx context.Context, in *pb.GetCheckpointsRequest) (*pb.GetCheckpointsResponse, error) {

}
