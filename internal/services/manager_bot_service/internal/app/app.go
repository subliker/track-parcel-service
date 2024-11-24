package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pm"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot"
)

type App interface {
	Run(context.Context) error
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

	a.logger.Info("app was built")
	return &a
}

func (a *app) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// receive sys signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.bot.Run(ctx); err != nil {
			a.logger.Errorf("bot stopped with error: %s", err)
			cancel()
		}
	}()
	a.logger.Info("app is running")

	// wait until signal will come or context will end
	select {
	case <-quit:
		a.logger.Info("shutdown signal received")
	case <-ctx.Done():
		a.logger.Info("context canceled")
	}

	a.logger.Info("stopping all service")
	// context cancel => bot stopping
	cancel()
	// wait until bot will be stopped
	wg.Wait()

	// stop clients
	if err := a.managerClient.Close(); err != nil {
		a.logger.Warn(err)
	}
	if err := a.parcelsManagerClient.Close(); err != nil {
		a.logger.Warn(err)
	}

	a.logger.Info("service was gracefully shutdowned)")
	return nil
}
