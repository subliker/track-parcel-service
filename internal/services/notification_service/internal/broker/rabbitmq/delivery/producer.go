package delivery

import "github.com/streadway/amqp"

type Producer struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func New(ch *amqp.Channel) (*Producer, error) {
	var p Producer

	deliveryQueue, err := ch.QueueDeclare(
		"notification_delivery",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	p.q = deliveryQueue

	return &p, nil
}
