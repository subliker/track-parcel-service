package main

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pm"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/session/lru"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
)

func main() {
	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg := config.Get()

	// creating new manager service client
	managerClient, err := manager.New(context.Background(), logger, cfg.ManagerClient)
	if err != nil {
		logger.Fatal(err)
	}

	// creating new parcels manager service client
	parcelsManagerClient, err := pm.New(context.Background(), logger, cfg.ParcelsManagerClient)
	if err != nil {
		logger.Fatal(err)
	}

	// creating lru session store
	store := lru.New(logger)

	// creating new bot
	bot := bot.New(logger, bot.BotOptions{
		Cfg:                  cfg.Bot,
		SessionStore:         store,
		ManagerClient:        managerClient,
		ParcelsManagerClient: parcelsManagerClient,
	})

	// creating new instance of app
	a := app.New(logger, app.AppOptions{
		Bot:                  bot,
		ManagerClient:        managerClient,
		ParcelsManagerClient: parcelsManagerClient,
	})
	// running app
	a.Run(context.Background())
}
