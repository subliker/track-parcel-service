package pu

import (
	"context"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *client) GetParcel(ctx context.Context, in *pb.GetParcelRequest) (*pb.GetParcelResponse, error) {
	logger := c.logger.WithFields("request", "get parcel")
	const errMsg = "request get parcel error: %s"

	// api call
	res, err := c.api.GetParcel(ctx, in)
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
