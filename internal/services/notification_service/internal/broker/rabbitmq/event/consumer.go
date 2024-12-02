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

	// setting logger
	c.logger = logger.WithFields("layer", "event consumer")

	// setting channel
	c.ch = ch

	// queue declaring
	eventsQueue, err := c.ch.QueueDeclare(
		"notification_events",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	c.q = eventsQueue

	// getting consumer channel
	eventsMsgs, err := ch.Consume(
		eventsQueue.Name, "",
		false, false, false,
		false, nil,
	)
	c.msgs = eventsMsgs

	// start messages receiving
	c.events = make(chan *notificationpb.Event)
	go c.receive()

	c.logger.Info("event consumer was successfully created")
	return &c, nil
}

func (c *consumer) Listen() <-chan *notificationpb.Event {
	return c.events
}

func (c *consumer) receive() {
	c.logger.Info("receiving messages running...")
	for msg := range c.msgs {
		event := notificationpb.Event{}

		// deserialization
		err := proto.Unmarshal(msg.Body, &event)
		if err != nil {
			errMsg := fmt.Errorf("error proto message deserialization: %s", err)
			c.logger.Error(errMsg)
			msg.Nack(false, false)
			continue
		}

		msg.Ack(false)
		c.events <- &event
	}
	c.logger.Info("receiving messages stopped")
}
