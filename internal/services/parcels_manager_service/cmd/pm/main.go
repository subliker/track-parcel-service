package main

import (
	"context"
	"flag"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/grpc"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/rest/api"
)

func main() {
	flag.Parse()

	// reading cfg
	cfg := config.Get()

	// creating logger
	logger := zap.NewLogger(cfg.Logger).WithFields("service", "parcels_manager_service")

	// creating store
	store, err := pg.New(logger, cfg.DB)
	if err != nil {
		logger.Fatalf("error store create: %s", err)
	}

	// creating broker
	broker, err := rabbitmq.New(logger, cfg.RabbitMQ)
	if err != nil {
		logger.Fatalf("error broker creating: %s", err)
	}

	// making event producer
	eventProducer, err := event.NewProducer(logger, broker.Chan())
	if err != nil {
		logger.Fatalf("error event producer making: %s", err)
	}

	// app context
	ctx := context.Background()

	// creating new manager client
	managerClient, err := manager.New(ctx, logger, cfg.ManagerClient)
	if err != nil {
		logger.Fatalf("error manager client creating: %s", err)
	}

	// creating new rest api server
	apiServer := api.New(logger, cfg.REST, managerClient, store, eventProducer)

	// creating new grpc server
	parcelServer := grpc.NewServer(logger, store, eventProducer)

	// creating new instance of app
	a := app.New(logger, app.AppOptions{
		Config:       cfg,
		Store:        store,
		ParcelServer: parcelServer,
		Broker:       broker,
		APIServer:    apiServer,
	})
	// running app
	a.Run(ctx)
}
