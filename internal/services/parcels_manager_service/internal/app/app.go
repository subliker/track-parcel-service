package app

import (
	"fmt"
	"net"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
	"google.golang.org/grpc"
)

type App interface {
	Run() error
}

type app struct {
	parcelServer *grpc.Server
	grpcConfig   config.GRPCConfig

	store parcel.Store

	logger logger.Logger
}

func New(cfg config.Config,
	logger logger.Logger,
	store parcel.Store,
	parcelServer *grpc.Server) App {
	var a app

	// set config
	a.grpcConfig = cfg.GRPC

	// set logger
	a.logger = logger.WithFields("layer", "app")

	// set store
	a.store = store

	// set parcel server
	a.parcelServer = parcelServer
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
	a.logger.Fatal(a.parcelServer.Serve(lis))
	return nil
}
