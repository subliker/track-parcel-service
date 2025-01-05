package main

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/delivery"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/dispatcher"
)

func main() {
	// reading config
	cfg := config.Get()

	// creating logger
	logger := zap.NewLogger(cfg.Logger).WithFields("service", "notification_service")

	// creating parcel store
	store, err := pg.New(logger, cfg.DB)
	if err != nil {
		logger.Fatal(err)
	}

	// creating broker
	broker, err := rabbitmq.New(logger, cfg.RabbitMQ)
	if err != nil {
		logger.Fatal(err)
	}

	// app context
	ctx := context.Background()

	// making event consumer
	eventConsumer, err := event.NewConsumer(ctx, logger, broker.Chan())
	if err != nil {
		logger.Fatal(err)
	}

	// making delivery producer
	deliverProducer, err := delivery.NewProducer(logger, broker.Chan())
	if err != nil {
		logger.Fatal(err)
	}

	// making dispatcher with producer and consumer
	dispatcher := dispatcher.New(logger, eventConsumer, deliverProducer, store)

	// creating new instance of app
	app := app.New(logger, app.AppOptions{
		Store:      store,
		Dispatcher: dispatcher,
		Broker:     broker,
	})
	// running app
	app.Run(ctx)
}
