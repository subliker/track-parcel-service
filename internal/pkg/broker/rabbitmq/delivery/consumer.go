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

func NewConsumer(ch *amqp.Channel) (Consumer, error) {
	var c consumer

	deliveryQueue, err := ch.QueueDeclare(
		"notification_delivery",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	c.q = deliveryQueue

	deliveryMsgs, err := ch.Consume(
		deliveryQueue.Name, "",
		false, false, false,
		false, nil,
	)
	c.msgs = deliveryMsgs

	c.delivery = make(chan *notificationpb.Delivery)
	go c.receive()

	return &c, nil
}

func (c *consumer) Listen() <-chan *notificationpb.Delivery {
	return c.delivery
}

func (c *consumer) receive() {
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
}
