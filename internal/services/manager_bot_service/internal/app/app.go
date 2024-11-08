package app

import (
	"context"

	sso "github.com/subliker/track-parcel-service/internal/pkg/client/sso/grpc"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session"
)

type App struct {
	bot       bot.Bot
	ssoClient sso.Client
	logger    logger.Logger
}

func New(cfg config.Config, logger logger.Logger) App {
	var a App

	// creation new session store
	store := session.New(logger)

	// creating new sso client
	ctx := context.Background()
	a.ssoClient = sso.New(ctx, logger, cfg.SSO)

	// creating new bot
	a.bot = bot.New(cfg.Bot, store, logger, a.ssoClient)

	// set logger
	a.logger = logger.WithFields("layer", "app")

	return a
}

func (a *App) Run() {
	// running bot
	a.logger.Fatal(a.bot.Run())
}
