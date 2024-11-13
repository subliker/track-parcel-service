package main

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/config"
)

func main() {
	// creating logger
	logger := zap.NewLogger()

	// reading cfg
	cfg := config.Get()

	// creating new instance of app
	a := app.New(cfg)
	// running app
}
