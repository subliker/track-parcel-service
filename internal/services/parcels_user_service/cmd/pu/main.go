package main

import (
	"context"
	"flag"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/server/grpc"
)

func main() {
	flag.Parse()

	// creating logger
	logger := zap.NewLogger()

	// reading config
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
	a.Run(context.Background())
}
