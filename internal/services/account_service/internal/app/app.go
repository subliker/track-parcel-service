package app

import (
	"fmt"
	"net"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc"
)

type App interface {
	Run() error
}

type app struct {
	accountServer *grpc.Server
	grpcConfig    config.GRPCConfig

	store store.Store

	logger logger.Logger
}

func New(cfg config.Config,
	logger logger.Logger,
	store store.Store,
	accountServer *grpc.Server) App {
	var a app

	// setting config
	a.grpcConfig = cfg.GRPC

	// setting store
	a.store = store

	// setting account server
	a.accountServer = accountServer

	// setting logger
	a.logger = logger.WithFields("layer", "app")
	return &a
}

func (a *app) Run() error {
	// creating new new listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcConfig.Port))
	if err != nil {
		a.logger.Fatal(err)
	}

	// running server
	a.logger.Infof("starting grpc server at port %d...", a.grpcConfig.Port)
	a.logger.Fatal(a.accountServer.Serve(lis))
	return nil
}
