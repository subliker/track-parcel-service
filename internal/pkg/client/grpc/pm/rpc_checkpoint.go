package pm

import (
	"context"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pmpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *client) AddCheckpoint(ctx context.Context, in *pb.AddCheckpointRequest) error {
	logger := c.logger.WithFields("request", "add checkpoint")
	const errMsg = "request add checkpoint error: %s"

	// api call
	_, err := c.api.AddCheckpoint(ctx, in)
	if err == nil {
		return nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return ErrParcelNotFound
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
func (c *client) GetCheckpoints(ctx context.Context, in *pb.GetCheckpointsRequest) (*pb.GetCheckpointsResponse, error) {
	logger := c.logger.WithFields("request", "get checkpoints")
	const errMsg = "request get checkpoints error: %s"

	// api call
	res, err := c.api.GetCheckpoints(ctx, in)
	if err == nil {
		return res, nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return nil, ErrParcelNotFound
		case codes.Internal:
			logger.Error(errMsg)
			return nil, ErrInternal
		default:
			logger.Error(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Errorf(errMsg, err)
	return nil, ErrUnexpected
}
