package pu

import (
	"context"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *client) AddSubscription(ctx context.Context, in *pb.AddSubscriptionRequest) error {
	logger := c.logger.WithFields("request", "add subscription")
	const errMsg = "request add subscription error: %s"

	// api call
	_, err := c.api.AddSubscription(ctx, in)
	if err == nil {
		return nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.AlreadyExists:
			return ErrAlreadyExists
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

func (c *client) DeleteSubscription(ctx context.Context, in *pb.DeleteSubscriptionRequest) error {
	logger := c.logger.WithFields("request", "delete subscription")
	const errMsg = "request delete subscription error: %s"

	// api call
	_, err := c.api.DeleteSubscription(ctx, in)
	if err == nil {
		return nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return ErrSubscriptionNotFound
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
