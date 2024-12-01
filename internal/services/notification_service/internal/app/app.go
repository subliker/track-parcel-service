package app

import (
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

func New(broker rabbitmq.Broker, dispatcher dispatcher.Notification) App {
	var a app

	a.broker = broker

	a.dispatcher = dispatcher

	return &a
}

func (a *app) Run() error {
	return nil
}

func (a *app) Close() error {
	return nil
}
