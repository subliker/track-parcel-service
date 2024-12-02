package app

import (
	"fmt"
	"net"

	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
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

	store parcel.ManagerStore

	broker rabbitmq.Broker

	logger logger.Logger
}

func New(cfg config.Config,
	logger logger.Logger,
	store parcel.ManagerStore,
	parcelServer *grpc.Server,
	broker rabbitmq.Broker) App {
	var a app

	// setting config
	a.grpcConfig = cfg.GRPC

	// setting logger
	a.logger = logger.WithFields("layer", "app")

	// setting store
	a.store = store

	// setting parcel server
	a.parcelServer = parcelServer

	// setting broker
	a.broker = broker
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
