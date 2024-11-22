package bot

import (
	"errors"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot/middleware"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot/style"
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
		Token:   cfg.Token,
		Poller:  &tele.LongPoller{Timeout: time.Second * 10},
		OnError: b.OnError,
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
	// global middlewares:
	// ensure sessions
	b.client.Use(middleware.Session(b.logger, b.sessionStore))

	// global handlers
	b.client.Handle("/start", b.handleStart())
	b.client.Handle("/add-parcel", b.handleAddParcel())
	b.client.Handle(&style.MenuBtnAddParcel, b.handleAddParcel())
	b.client.Handle(tele.OnText, b.handleOnText())

	// groups
	// group for auth middleware
	authGroup := b.client.Group()
	authGroup.Use(middleware.Auth(b.logger, b.managerClient))
	// handle register
	authGroup.Handle("/register", b.handleRegister())
	authGroup.Handle(&style.MenuBtnRegister, b.handleRegister())
}

func (b *bot) OnError(err error, ctx tele.Context) {
	logger := b.logger.WithFields("user_id", ctx.Sender().ID)

	// if ctx is nil
	if ctx == nil {
		logger.Error("handler ended with empty context and error: %s", err)
	}

	// add handler name
	handlerInterface := ctx.Get("handler")
	handler, ok := handlerInterface.(string)
	if ok {
		logger.WithFields("handler", handler)
	}

	// add state handler name
	stateHandlerInterface := ctx.Get("state_handler")
	stateHandler, ok := stateHandlerInterface.(string)
	if ok {
		logger.WithFields("state_handler", stateHandler)
	}

	// log error
	logger.Errorf("handler ended with error: %s", err)
}
