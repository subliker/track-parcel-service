package main

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/lru"
)

func main() {
	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg := config.Get()

	// creating new manager service client
	managerClient := manager.New(context.Background(), logger, cfg.ManagerService)

	// creating lru session store
	store := lru.New(logger)

	// creating new bot
	bot := bot.New(cfg.Bot, store, logger, managerClient)

	// creating new instance of app
	a := app.New(cfg, logger, bot, managerClient)
	// running app
	a.Run()
}
