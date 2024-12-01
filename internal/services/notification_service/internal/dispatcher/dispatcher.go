package dispatcher

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/delivery"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/store/parcel"
)

type Notification interface {
	Run() error
}

type notification struct {
	eventConsumer    event.Consumer
	deliveryProducer delivery.Producer
	store            parcel.NotificationStore

	logger logger.Logger
}

func New(
	logger logger.Logger,
	eventConsumer event.Consumer,
	deliveryProducer delivery.Producer,
	store parcel.NotificationStore,
) Notification {
	var n notification

	n.eventConsumer = eventConsumer
	n.deliveryProducer = deliveryProducer

	n.store = store

	n.logger = logger
	return &n
}

func (n *notification) Run() error {
	return nil
}
