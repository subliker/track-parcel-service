package manager

import (
	"context"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *client) Register(ctx context.Context, in *pb.RegisterRequest) (*emptypb.Empty, error) {
	logger := c.logger.WithFields("request", "register")
	const errMsg = "request register error: %s"

	// api call
	res, err := c.api.Register(context.Background(), in)
	if err == nil {
		return res, nil
	}

	// handle errors
	if grpcStatus, ok := status.FromError(err); ok {
		errMsg := fmt.Errorf(errMsg, grpcStatus.Message())

		switch grpcStatus.Code() {
		case codes.AlreadyExists:
			return nil, ErrManagerIsAlreadyExist
		case codes.Internal:
			logger.Info(errMsg)
			return nil, ErrInternal
		default:
			logger.Info(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Infof(errMsg, "non grpc error")
	return nil, ErrUnexpected
}

func (c *client) GetInfo(ctx context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	logger := c.logger.WithFields("request", "register")
	const errMsg = "request register error: %s"

	// api call
	res, err := c.api.GetInfo(context.Background(), in)
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
			logger.Info(errMsg)
			return nil, ErrInternal
		default:
			logger.Info(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Infof(errMsg, "non grpc error")
	return nil, ErrUnexpected
}

func (c *client) GetApiToken(ctx context.Context, in *pb.GetApiTokenRequest) (*pb.GetApiTokenResponse, error) {
	logger := c.logger.WithFields("request", "register")
	const errMsg = "request register error: %s"

	// api call
	res, err := c.api.GetApiToken(context.Background(), in)
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
			logger.Info(errMsg)
			return nil, ErrInternal
		default:
			logger.Info(errMsg)
			return nil, ErrUnexpected
		}
	}
	logger.Infof(errMsg, "non grpc error")
	return nil, ErrUnexpected
}
