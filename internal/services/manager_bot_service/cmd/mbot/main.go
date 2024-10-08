package main

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/app"
)

func main() {
	if err := app.New(); err != nil {
		logger.Zap.Fatalf("error running app: %s", err)
	}
}
