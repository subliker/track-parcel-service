package main

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq/delivery"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pu"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/session/lru"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/bot"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/config"
)

func main() {
	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg := config.Get()

	// creating new user service client
	userClient, err := user.New(context.Background(), logger, cfg.UserClient)
	if err != nil {
		logger.Fatal(err)
	}

	// creating new parcels user service client
	parcelsUserClient, err := pu.New(context.Background(), logger, cfg.ParcelsUserClient)
	if err != nil {
		logger.Fatal(err)
	}

	// creating lru session store
	store := lru.New(logger)

	// creating broker
	broker, err := rabbitmq.New(logger, cfg.RabbitMQ)
	if err != nil {
		logger.Fatal(err)
	}

	// creating delivery consumer
	deliveryConsumer, err := delivery.NewConsumer(logger, broker.Chan())
	if err != nil {
		logger.Fatal(err)
	}

	// creating new bot
	bot := bot.New(logger, bot.BotOptions{
		Cfg:               cfg.Bot,
		SessionStore:      store,
		UserClient:        userClient,
		ParcelsUserClient: parcelsUserClient,
		DeliveryConsumer:  deliveryConsumer,
	})

	// creating new instance of app
	a := app.New(logger, app.AppOptions{
		Bot:               bot,
		UserClient:        userClient,
		ParcelsUserClient: parcelsUserClient,
		Broker:            broker,
	})
	// running app
	a.Run(context.Background())
}
