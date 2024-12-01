package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/delivery"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/config"
)

type Broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel

	deliveryProducer delivery.Producer
	eventsConsumer   event.Consumer
}

func New(logger logger.Logger, cfg config.RabbitMQConfig) (*Broker, error) {
	var b Broker

	// open mq connection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", cfg.User, cfg.Password, cfg.Host))
	if err != nil {
		return nil, err
	}
	b.conn = conn

	// getting channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	b.ch = ch

	// make producers and consumers
	deliveryProducer, err := delivery.New(logger, ch)
	if err != nil {
		return nil, fmt.Errorf("error making delivery producer: %s", err)
	}
	b.deliveryProducer = deliveryProducer

	eventsConsumer, err := event.New(ch)
	if err != nil {
		return nil, fmt.Errorf("error making events consumer: %s", err)
	}
	b.eventsConsumer = eventsConsumer

	return &b, nil
}
