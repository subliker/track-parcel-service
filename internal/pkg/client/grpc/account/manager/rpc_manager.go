package manager

import (
	"context"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *client) Register(ctx context.Context, in *pb.RegisterRequest) error {
	logger := c.logger.WithFields("request", "register")
	const errMsg = "request register error: %s"

	// api call
	_, err := c.api.Register(ctx, in)
	if err == nil {
		return nil
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

func (c *client) GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	logger := c.logger.WithFields("request", "get info")
	const errMsg = "request get info error: %s"

	// api call
	res, err := c.api.GetInfo(ctx, in)
	if err == nil {
		return res, nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return nil, ErrManagerNotFound
		case codes.Internal:
			logger.Error(errMsg)
			return nil, ErrInternal
		default:
			logger.Error(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Errorf(errMsg, "non grpc error")
	return nil, ErrUnexpected
}

func (c *client) GetApiToken(ctx context.Context, in *pb.GetApiTokenRequest) (*pb.GetApiTokenResponse, error) {
	logger := c.logger.WithFields("request", "get api token")
	const errMsg = "request get api token error: %s"

	// api call
	res, err := c.api.GetApiToken(ctx, in)
	if err == nil {
		return res, nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return nil, ErrManagerNotFound
		case codes.Internal:
			logger.Error(errMsg)
			return nil, ErrInternal
		default:
			logger.Error(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Errorf(errMsg, "non grpc error")
	return nil, ErrUnexpected
}

func (c *client) Auth(ctx context.Context, in *pb.AuthRequest) error {
	logger := c.logger.WithFields("request", "auth")
	const errMsg = "request auth error: %s"

	// api call
	_, err := c.api.Auth(ctx, in)
	if err == nil {
		return nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return ErrManagerNotFound
		case codes.Internal:
			logger.Error(errMsg)
			return ErrInternal
		default:
			logger.Error(errMsg)
			return ErrUnexpected
		}
	}
	logger.Errorf(errMsg, "non grpc error")
	return ErrUnexpected
}

func (c *client) AuthApiToken(ctx context.Context, in *pb.AuthApiTokenRequest) (*pb.AuthApiTokenResponse, error) {
	logger := c.logger.WithFields("request", "auth api token")
	const errMsg = "request auth api token error: %s"

	// api call
	res, err := c.api.AuthApiToken(ctx, in)
	if err == nil {
		return res, nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.NotFound:
			return nil, ErrManagerNotFound
		case codes.Internal:
			logger.Error(errMsg)
			return nil, ErrInternal
		default:
			logger.Error(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Errorf(errMsg, "non grpc error")
	return nil, ErrUnexpected
}
