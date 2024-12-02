package event

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"google.golang.org/protobuf/proto"
)

type Producer interface {
	Publish(*notificationpb.Event) error
}

type producer struct {
	ch *amqp.Channel
	q  amqp.Queue

	logger logger.Logger
}

func NewProducer(logger logger.Logger, ch *amqp.Channel) (Producer, error) {
	var p producer

	// setting logger
	p.logger = logger.WithFields("layer", "event producer")

	// setting channel
	p.ch = ch

	// queue declaring
	deliveryQueue, err := p.ch.QueueDeclare(
		"notification_events",
		true, false, false,
		false, nil,
	)
	if err != nil {
		return nil, err
	}
	p.q = deliveryQueue

	p.logger.Info("event producer was successfully created")
	return &p, nil
}

func (p *producer) Publish(event *notificationpb.Event) error {
	// protobuf serialization
	body, err := proto.Marshal(event)
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
