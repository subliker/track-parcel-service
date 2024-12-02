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

	logger logger.Logger
}

func New(logger logger.Logger, cfg Config) (Broker, error) {
	var b broker

	// setting logger
	b.logger = logger.WithFields("layer", "broker")

	// open mq connection
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", cfg.User, cfg.Password, cfg.Host))
	if err != nil {
		return nil, err
	}
	b.conn = conn
	b.logger.Infof("broker was connected to %s", conn.LocalAddr())

	// getting channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	b.ch = ch

	b.logger.Info("broker was successfully created")
	return &b, nil
}

func (b *broker) Chan() *amqp.Channel {
	return b.ch
}

func (b *broker) Close() error {
	// connection close
	if err := b.conn.Close(); err != nil {
		b.logger.Warnf("broker connection closed with error: %s", err)
	}

	b.logger.Info("broker was closed")
	return nil
}
