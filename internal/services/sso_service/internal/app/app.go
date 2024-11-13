package app

import (
	"net"

	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/grpc/auth"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store/pgstore"
	"google.golang.org/grpc"
)

type app struct {
	ssoServer *grpc.Server

	logger logger.Logger
}

type App interface {
	Run() error
}

func New(cfg config.Config, logger logger.Logger) App {
	var a app

	// set logger
	a.logger = logger.WithFields("layer", "app")

	// setting up store
	store, err := pgstore.New(cfg.DB)
	if err != nil {
		zap.Logger.Fatal(err)
	}

	// creating new sso server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		zap.Logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	ssov1.RegisterAuthServer(grpcServer, auth.New(store))

	a.ssoServer = grpcServer
	return &a
}

func (a *app) Run() error {
	// running server
	if err := a.ssoServer.Serve(lis); err != nil {
		return err
	}

}
