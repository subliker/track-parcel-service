package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/dispatcher"
)

// App is interface to control app.
// Closing app closes all app services
type App interface {
	// Run starts app and stops all internal services when ctx done.
	Run(context.Context) error
}
type app struct {
	broker rabbitmq.Broker

	dispatcher dispatcher.Notification

	logger logger.Logger
}

// New creates new app instance
func New(
	logger logger.Logger,
	broker rabbitmq.Broker,
	dispatcher dispatcher.Notification,
) App {
	var a app

	// setting logger
	a.logger = logger.WithFields("layer", "app")

	// setting broker
	a.broker = broker

	// setting dispatcher
	a.dispatcher = dispatcher

	a.logger.Info("app was created")
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
		// starting dispatcher
		if err := a.dispatcher.Run(ctx); err != nil {
			errMsg := fmt.Errorf("running dispatcher error: %s", err)
			a.logger.Error(errMsg)
			cancel()
		}
	}()
	a.logger.Info("app running...")

	// wait until signal will come or context will end
	select {
	case <-quit:
		a.logger.Info("shutdown signal received")
	case <-ctx.Done():
		a.logger.Info("context canceled")
	}

	a.logger.Info("stopping all services")
	// stop services
	cancel()
	// wait until services will be stopped
	wg.Wait()

	// stop broker
	if err := a.broker.Close(); err != nil {
		a.logger.Warn(err)
	}

	a.logger.Info("app was gracefully shutdowned :)")
	return nil
}
