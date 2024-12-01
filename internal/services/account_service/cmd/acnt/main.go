package main

import (
	"flag"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/server/grpc"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store/pg"
)

func main() {
	flag.Parse()

	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg := config.Get()

	// creating parcel store
	store, err := pg.New(logger, cfg.DB)
	if err != nil {
		logger.Fatalf("error store create: %s", err)
	}

	// creating new grpc server
	accountServer := grpc.NewServer(logger, store)

	// creating new instance of app
	a := app.New(cfg, logger, store, accountServer)
	// running app
	a.Run()
}
