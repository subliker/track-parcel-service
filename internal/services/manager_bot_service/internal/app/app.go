package app

import (
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
)

type app struct{}

func New() error {
	cfg := config.Get()
	bot.Run(bot.BotConfig{
		Token: cfg.Token,
	})
	return nil
}
