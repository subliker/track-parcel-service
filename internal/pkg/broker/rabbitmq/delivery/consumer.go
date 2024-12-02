package delivery

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"google.golang.org/protobuf/proto"
)

type Consumer interface {
	Listen() <-chan *notificationpb.Delivery
}

type consumer struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
	q    amqp.Queue

	delivery chan *notificationpb.Delivery

	logger logger.Logger
}

func NewConsumer(logger logger.Logger, ch *amqp.Channel) (Consumer, error) {
	var c consumer

	// setting logger
	c.logger = logger.WithFields("layer", "delivery consumer")

	// setting channel
	c.ch = ch

	// queue declaring
	deliveryQueue, err := c.ch.QueueDeclare(
		"notification_delivery",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	c.q = deliveryQueue

	// getting consumer channel
	deliveryMsgs, err := ch.Consume(
		deliveryQueue.Name, "",
		false, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	c.msgs = deliveryMsgs

	// start messages receiving
	c.delivery = make(chan *notificationpb.Delivery)
	go c.receive()

	c.logger.Info("delivery consumer was successfully created")
	return &c, nil
}

func (c *consumer) Listen() <-chan *notificationpb.Delivery {
	return c.delivery
}

func (c *consumer) receive() {
	c.logger.Info("receiving messages running...")
	for msg := range c.msgs {
		event := notificationpb.Delivery{}

		// deserialization
		err := proto.Unmarshal(msg.Body, &event)
		if err != nil {
			errMsg := fmt.Errorf("error proto message deserialization: %s", err)
			c.logger.Error(errMsg)
			msg.Nack(false, true)
			continue
		}

		c.delivery <- &event
		msg.Ack(false)
	}
	c.logger.Info("receiving messages stopped")
}
