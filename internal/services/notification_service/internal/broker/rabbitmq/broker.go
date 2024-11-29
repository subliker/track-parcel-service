package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/delivery"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/events"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/config"
)

type Broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel

	deliveryProducer delivery.Producer
	eventsConsumer   events.Consumer
}

func New(cfg config.RabbitMQConfig) (*Broker, error) {
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
	deliveryProducer, err := delivery.New(ch)
	if err != nil {
		return nil, fmt.Errorf("error making delivery producer: %s", err)
	}
	b.deliveryProducer = deliveryProducer

	eventsConsumer, err := events.New(ch)
	if err != nil {
		return nil, fmt.Errorf("error making events consumer: %s", err)
	}
	b.eventsConsumer = eventsConsumer

	return &b, nil
}
