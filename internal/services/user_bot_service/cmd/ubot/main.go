package main

import (
	"context"
	"flag"

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

var ignoreBroker *bool

func main() {
	// flags
	ignoreBroker = flag.Bool("ignore-broker", false, "ignore broker connection")
	flag.Parse()

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

	// creating broker and delivery consumer
	var broker rabbitmq.Broker
	var deliveryConsumer delivery.Consumer
	if !*ignoreBroker {
		// creating broker
		broker, err := rabbitmq.New(logger, cfg.RabbitMQ)
		if err != nil {
			logger.Fatal(err)
		}
		// creating delivery consumer
		deliveryConsumer, err = delivery.NewConsumer(logger, broker.Chan())
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Warn("using broker was ignored")
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
