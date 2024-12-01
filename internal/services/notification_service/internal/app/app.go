package app

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/dispatcher"
)

type App interface {
	Run() error
	Close() error
}
type app struct {
	broker     rabbitmq.Broker
	dispatcher dispatcher.Notification

	logger logger.Logger
}

func New(logger logger.Logger, broker rabbitmq.Broker, dispatcher dispatcher.Notification) App {
	var a app

	// setting broker
	a.broker = broker

	// setting dispatcher
	a.dispatcher = dispatcher

	// setting logger
	a.logger = logger.WithFields("layer", "app")

	return &a
}

func (a *app) Run() error {
	a.logger.Info("app running...")

	// starting dispatcher
	if err := a.dispatcher.Run(); err != nil {
		errMsg := fmt.Errorf("running dispatcher error: %s", err)
		a.logger.Error(errMsg)
		return errMsg
	}

	a.logger.Info("app stopped")
	return nil
}

func (a *app) Close() error {
	return nil
}
