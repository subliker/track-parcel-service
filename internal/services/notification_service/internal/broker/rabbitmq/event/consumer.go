package event

import "github.com/streadway/amqp"

type Consumer interface {
	Listen()
}

type consumer struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
	q    amqp.Queue
}

func New(ch *amqp.Channel) (*consumer, error) {
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

	return &c, nil
}

func (c *consumer) Listen() {}
