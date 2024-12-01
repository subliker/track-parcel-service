package event

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"google.golang.org/protobuf/proto"
)

type Consumer interface {
	Listen() <-chan *notificationpb.Event
}

type consumer struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
	q    amqp.Queue

	events chan *notificationpb.Event

	logger logger.Logger
}

func NewConsumer(logger logger.Logger, ch *amqp.Channel) (Consumer, error) {
	var c consumer

	eventsQueue, err := ch.QueueDeclare(
		"notification_events",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	c.q = eventsQueue

	eventsMsgs, err := ch.Consume(
		eventsQueue.Name, "",
		false, false, false,
		false, nil,
	)
	c.msgs = eventsMsgs

	c.events = make(chan *notificationpb.Event)
	go c.receive()

	return &c, nil
}

func (c *consumer) Listen() <-chan *notificationpb.Event {
	return c.events
}

func (c *consumer) receive() {
	for msg := range c.msgs {
		event := notificationpb.Event{}

		// deserialization
		err := proto.Unmarshal(msg.Body, &event)
		if err != nil {
			errMsg := fmt.Errorf("error proto message deserialization: %s", err)
			c.logger.Error(errMsg)
			continue
		}

		c.events <- &event
	}
}
