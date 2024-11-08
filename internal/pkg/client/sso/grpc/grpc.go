package sso

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	ssov1 "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	api    ssov1.AuthClient
	logger logger.Logger
}

type Client interface {
	RegisterTelegramID(ctx context.Context, in *ssov1.RegisterTelegramIDRequest) (*ssov1.RegisterTelegramIDResponse, error)
}

func New(ctx context.Context, logger logger.Logger, cfg Config) Client {
	cc, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(cc.Target())

	return &client{
		api:    ssov1.NewAuthClient(cc),
		logger: logger.WithFields("layer", "grpc client", "service", "sso"),
	}
}

func (c *client) RegisterTelegramID(ctx context.Context, in *ssov1.RegisterTelegramIDRequest) (*ssov1.RegisterTelegramIDResponse, error) {
	const requestName = "register telegram id"
	const errMsg = "request " + requestName + " error: %s"
	logger := c.logger.WithFields("request", requestName)

	res, err := c.api.RegisterTelegramID(ctx, in)
	if err != nil {
		logger.Errorf(errMsg, err)
		return nil, fmt.Errorf(errMsg, err)
	}

	return res, nil
}
