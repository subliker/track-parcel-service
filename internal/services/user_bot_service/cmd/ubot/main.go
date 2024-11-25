package main

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/session/lru"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/config"
)

func main() {
	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg := config.Get()

	// creating new user service client
	userClient, err := user.New(context.Background(), logger, cfg.UserClient)
	if err != nil {
		logger.Fatal(err)
	}

	// creating new parcels user service client
	//

	// creating lru session store
	store := lru.New(logger)

	// creating new bot
	bot := bot.New(logger, bot.BotOptions{
		Cfg:          cfg.Bot,
		SessionStore: store,
		UserClient:   userClient,
	})

	// creating new instance of app
	a := app.New(logger, app.AppOptions{
		Bot:        bot,
		UserClient: userClient,
	})
	// running app
	a.Run(context.Background())
}
