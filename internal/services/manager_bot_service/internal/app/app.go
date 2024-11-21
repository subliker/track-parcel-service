package app

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session"
)

type App struct {
	bot       bot.Bot
	ssoClient manager.Client

	logger logger.Logger
}

func New(cfg config.Config, logger logger.Logger) App {
	var a App

	// set logger
	a.logger = logger.WithFields("layer", "app")

	// creation new session store
	store := session.New(logger)

	// creating new bot
	a.bot = bot.New(cfg.Bot, store, logger, a.ssoClient)

	// creating new sso client
	ctx := context.Background()
	a.ssoClient = manager.New(ctx, logger, cfg.SSO)

	return a
}

func (a *App) Run() {
	// running bot
	a.logger.Fatal(a.bot.Run())
}
