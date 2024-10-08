package main

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		logger.Zap.Fatal(err)
	}
}
