package dispatcher

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/delivery"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/notification_service/internal/store/parcel"
)

// Notification is dispatcher, that receives events and sends notifications(delivery).
type Notification interface {
	// Run runs notification dispatcher. Stops when ctx done.
	Run(context.Context) error
}

type notification struct {
	eventConsumer    event.Consumer
	deliveryProducer delivery.Producer

	store parcel.NotificationStore

	logger logger.Logger
}

// New creates new instance of notification dispatcher
func New(
	logger logger.Logger,
	eventConsumer event.Consumer,
	deliveryProducer delivery.Producer,
	store parcel.NotificationStore,
) Notification {
	var n notification

	// setting producers and consumers
	n.eventConsumer = eventConsumer
	n.deliveryProducer = deliveryProducer

	// setting store
	n.store = store

	n.logger = logger.WithFields("layer", "notification dispatcher")
	return &n
}

// Run runs notification dispatcher
func (n *notification) Run(ctx context.Context) error {
	n.logger.Info("notification dispatcher running...")

	// handling events until ctx done
	n.handleEvents(ctx)

	// close consumers
	n.eventConsumer.Close()

	n.logger.Info("notification dispatcher stopped")
	return nil
}
