package app

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/store/session"
)

type App struct {
	bot bot.Bot
}

func New(cfg config.Config) App {
	var a App

	// creation new session store
	store := session.New()

	// creating new bot
	a.bot = bot.New(cfg.Bot, store)

	return a
}

func (a *App) Run() {
	// running bot
	logger.Zap.Fatal(a.bot.Run())
}
