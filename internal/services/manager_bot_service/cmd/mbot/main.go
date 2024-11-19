package main

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
)

func main() {
	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg := config.Get()

	// creating new instance of app
	a := app.New(cfg, logger)
	// running app
	a.Run()
}
