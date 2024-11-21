package app

import (
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
)

type App struct {
	bot           bot.Bot
	managerClient manager.Client

	logger logger.Logger
}

func New(cfg config.Config,
	logger logger.Logger,
	bot bot.Bot,
	managerClient manager.Client) App {
	var a App

	// set logger
	a.logger = logger.WithFields("layer", "app")

	// set bot
	a.bot = bot

	// set manager client
	a.managerClient = managerClient

	return a
}

func (a *App) Run() {
	// running bot
	a.logger.Fatal(a.bot.Run())
}
