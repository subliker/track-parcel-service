package bot

import (
	"errors"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot/middleware"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/lang"
	tele "gopkg.in/telebot.v4"
)

type Bot interface {
	Run() error
}

type bot struct {
	client        *tele.Bot
	bundle        lang.Messages
	sessionStore  session.Store
	managerClient manager.Client
	logger        logger.Logger
}

// New creates new instance of bot
func New(cfg config.BotConfig, ss session.Store, logger logger.Logger, managerClient manager.Client) Bot {
	var b bot

	// try to build bot client
	client, err := tele.NewBot(tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: time.Second * 10},
	})
	if err != nil {
		logger.Fatalf("error building bot: %s", err)
	}
	b.client = client

	// set sso client
	b.managerClient = managerClient

	// set session store
	b.sessionStore = ss

	// language initialization
	b.bundle = lang.MessagesForOrDefault(cfg.Language)

	// set logger
	b.logger = logger.WithFields("layer", "bot")

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
	// ensure session
	b.client.Use(middleware.Session(b.sessionStore))

	b.client.Handle("/start", b.handleStart())
	b.client.Handle("/add-parcel", b.handleAddParcel())
	b.client.Handle("/register", b.handleRegister())

	b.client.Handle(tele.OnText, b.handleOnText())
}
