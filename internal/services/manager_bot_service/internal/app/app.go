package app

import (
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pm"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
)

type App interface {
	Run()
}

type app struct {
	bot                  bot.Bot
	managerClient        manager.Client
	parcelsManagerClient pm.Client

	logger logger.Logger
}

type AppOptions struct {
	Bot                  bot.Bot
	ManagerClient        manager.Client
	ParcelsManagerClient pm.Client
}

func New(logger logger.Logger, opts AppOptions) App {
	var a app

	// set logger
	a.logger = logger.WithFields("layer", "app")

	// set bot
	a.bot = opts.Bot

	// set manager client
	a.managerClient = opts.ManagerClient

	// set parcels manager client
	a.parcelsManagerClient = opts.ParcelsManagerClient

	return &a
}

func (a *app) Run() {
	// running bot
	a.logger.Fatal(a.bot.Run())
}
