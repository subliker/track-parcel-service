package main

import (
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
	cfg := config.Get()

	logger := zap.NewLogger()

	store, err := pg.New(logger, cfg.DB)
	if err != nil {
		logger.Fatal(err)
	}

	broker, err := rabbitmq.New(logger, cfg.RabbitMQ)
	if err != nil {
		logger.Fatal(err)
	}

	eventConsumer, err := event.NewConsumer(logger, broker.Chan())
	if err != nil {
		logger.Fatal(err)
	}

	deliverProducer, err := delivery.NewProducer(logger, broker.Chan())
	if err != nil {
		logger.Fatal(err)
	}

	dispatcher := dispatcher.New(logger, eventConsumer, deliverProducer, store)

	app := app.New(broker, dispatcher)

	app.Run()
}
