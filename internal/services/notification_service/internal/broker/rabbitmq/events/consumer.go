package events

import "github.com/streadway/amqp"

type Consumer struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func New(ch *amqp.Channel) (*Consumer, error) {
	var c Consumer

	eventsQueue, err := ch.QueueDeclare(
		"notification_events",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	c.q = eventsQueue

	return &c, nil
}
