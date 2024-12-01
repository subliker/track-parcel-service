package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
)

type Broker interface {
	Chan() *amqp.Channel
	Close() error
}

type broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(logger logger.Logger, cfg Config) (Broker, error) {
	var b broker

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

	return &b, nil
}

func (b *broker) Chan() *amqp.Channel {
	return b.ch
}

func (b *broker) Close() error {
	return nil
}
