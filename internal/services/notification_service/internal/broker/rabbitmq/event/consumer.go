package event

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"google.golang.org/protobuf/proto"
)

// Consumer receives events and transfers into pb format
type Consumer interface {
	// Listen receives events transfered into pb
	Listen() <-chan *notificationpb.Event
	// Close closes listen channel
	Close()
}

type consumer struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
	q    amqp.Queue

	events chan *notificationpb.Event

	logger logger.Logger
}

// NewConsumer creates new instance of event consumer
func NewConsumer(ctx context.Context, logger logger.Logger, ch *amqp.Channel) (Consumer, error) {
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
	if err != nil {
		return nil, err
	}
	c.msgs = eventsMsgs

	// start messages receiving
	c.events = make(chan *notificationpb.Event)

	// consumer receives msgs until ctx done
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
		fmt.Print(msg)
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
}

func (c *consumer) Close() {
	// close events channel
	close(c.events)
	c.logger.Info("receiving messages stopped")
}
