package bot

import (
	"strings"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pm"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot/middleware"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/lang"
	"golang.org/x/net/context"
	tele "gopkg.in/telebot.v4"
)

type Bot interface {
	Run(context.Context) error
}

type bot struct {
	client       *tele.Bot
	bundle       lang.Messages
	sessionStore session.Store

	managerClient        manager.Client
	parcelsManagerClient pm.Client

	logger logger.Logger
}

type BotOptions struct {
	Cfg                  config.BotConfig
	SessionStore         session.Store
	ManagerClient        manager.Client
	ParcelsManagerClient pm.Client
}

// New creates new instance of bot
func New(logger logger.Logger, opts BotOptions) Bot {
	var b bot

	// try to build bot client
	client, err := tele.NewBot(tele.Settings{
		Token:     strings.TrimSpace(opts.Cfg.Token),
		Poller:    &tele.LongPoller{Timeout: time.Second * 10},
		OnError:   b.OnError,
		ParseMode: tele.ModeMarkdown,
	})
	if err != nil {
		logger.Fatalf("error building bot: %s", err)
	}
	b.client = client

	// set manager client
	b.managerClient = opts.ManagerClient

	// set parcels manager client
	b.parcelsManagerClient = opts.ParcelsManagerClient

	// set session store
	b.sessionStore = opts.SessionStore

	// language initialization
	b.bundle = lang.MessagesForOrDefault(opts.Cfg.Language)

	// set logger
	b.logger = logger.WithFields("layer", "bot")

	// handlers init
	b.initHandlers()

	b.logger.Infof("bot was built. Hello, I'm %s", b.client.Me.FirstName)
	return &b
}

// Run runs bot after initialization
func (b *bot) Run(ctx context.Context) error {
	go func() {
		b.client.Start()
	}()
	b.logger.Info("bot is running")

	// wait until context will be canceled
	<-ctx.Done()

	// stop bot
	b.client.Close()

	return nil
}

// initHandlers initializes all handlers
func (b *bot) initHandlers() {
	// global middlewares:
	// ensure sessions
	b.client.Use(middleware.Session(b.logger, b.sessionStore, b.bundle))
	b.client.Use(middleware.Auth(b.logger, b.managerClient, b.bundle))

	// global handlers
	// handle text
	b.client.Handle(tele.OnText, b.handleOnText())
	// handle start
	b.client.Handle("/start", b.handleStart())
	// handle register
	registerHandler := b.handleRegister()
	b.client.Handle("/register", registerHandler)
	b.client.Handle(&startBtnRegister, registerHandler)
	// don't specify data handler
	b.client.Handle(&btnNotSpecify, b.handleDontSpecify())

	// groups
	// group for authorized managers middleware
	authGroup := b.client.Group()
	authGroup.Use(middleware.Authorized(b.logger, b.bundle))
	// handle menu
	authGroup.Handle("/menu", b.handleMenu())
	// handle add parcel
	addParcelHandler := b.handleAddParcel()
	authGroup.Handle("/add-parcel", addParcelHandler)
	authGroup.Handle(&menuBtnAddParcel, addParcelHandler)
	// handle my api
	myApiHandler := b.handleMyApi()
	authGroup.Handle("/my-api", myApiHandler)
	authGroup.Handle(&menuBtnMyApi, myApiHandler)

	b.logger.Info("handlers were initialized")
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
