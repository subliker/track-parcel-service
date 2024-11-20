package main

import (
	"flag"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/grpc"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel/pg"
)

func main() {
	flag.Parse()

	// creating logger
	logger := zap.NewLogger()

	// reading cfg
	cfg := config.Get()

	// creating store
	store, err := pg.New(logger, cfg.DB)
	if err != nil {
		logger.Fatalf("error store create: %s", err)
	}

	// creating new grpc server
	parcelServer := grpc.NewServer(logger, store)

	// creating new instance of app
	a := app.New(cfg, logger, store, parcelServer)
	// running app
	a.Run()
}
