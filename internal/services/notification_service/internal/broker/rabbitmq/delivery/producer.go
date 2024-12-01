package delivery

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"google.golang.org/protobuf/proto"
)

type Producer interface {
	Publish(*notificationpb.Delivery) error
}

type producer struct {
	ch *amqp.Channel
	q  amqp.Queue

	logger logger.Logger
}

func NewProducer(logger logger.Logger, ch *amqp.Channel) (Producer, error) {
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

	p.logger = logger.WithFields("producer", "delivery")
	return &p, nil
}

func (p *producer) Publish(delivery *notificationpb.Delivery) error {
	// protobuf serialization
	body, err := proto.Marshal(delivery)
	if err != nil {
		errMsg := fmt.Errorf("error marshaling proto message: %s", err)
		p.logger.Error(errMsg)
		return errMsg
	}

	// publishing
	err = p.ch.Publish(
		"",
		p.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/x-protobuf",
			Body:        body,
		},
	)
	if err != nil {
		errMsg := fmt.Errorf("error publishing proto message: %s", err)
		p.logger.Error(errMsg)
		return errMsg
	}
	return nil
}
