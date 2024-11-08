package app

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session"
)

type App struct {
	bot    bot.Bot
	logger logger.Logger
}

func New(cfg config.Config, logger logger.Logger) App {
	var a App

	// creation new session store
	store := session.New(logger)

	// creating new bot
	a.bot = bot.New(cfg.Bot, store, logger)

	// set logger
	a.logger = logger.WithFields("layer", "app")

	return a
}

func (a *App) Run() {
	// running bot
	a.logger.Fatal(a.bot.Run())
}
