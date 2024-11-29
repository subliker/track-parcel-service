package delivery

import "github.com/streadway/amqp"

type Producer interface {
	Publish() error
}

type producer struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func New(ch *amqp.Channel) (Producer, error) {
	var p producer

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

func (p *producer) Publish() error {
	return nil
}
