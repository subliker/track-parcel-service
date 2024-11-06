package bot

import (
	"errors"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	tele "gopkg.in/telebot.v4"
)

type Bot interface {
	Run() error
}

type bot struct {
	client *tele.Bot
}

// New creates new instance of bot
func New(cfg config.BotConfig, ss session.Store) Bot {
	var b bot

	// try to build bot client
	client, err := tele.NewBot(tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: time.Second * 10},
	})
	if err != nil {
		logger.Zap.Fatalf("error building bot: %s", err)
	}
	b.client = client

	// handlers init
	b.initHandlers()

	return &b
}

// Run runs bot after initialization
func (b *bot) Run() error {
	b.client.Start()
	return errors.New("bot stopped")
}

// initHandlers initializes all handlers
func (b *bot) initHandlers() {
	b.client.Handle("/start", func(c tele.Context) error {
		return c.Reply("Hello!")
	})
}
