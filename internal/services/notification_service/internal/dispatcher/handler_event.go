package dispatcher

import (
	"context"

	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
)

// handleEvents catches event messages from event consumer and sends notifications to subscribers.
// Stops when ctx done.
func (n *notification) handleEvents(ctx context.Context) {
	logger := n.logger.WithFields("handler", "event")

	// listening events from consumer and wait until ctx done
	for {
		select {
		case event := <-n.eventConsumer.Listen():
			// getting parcel subscribers
			subs, err := n.store.ParcelSubscribers(model.TrackNumber(event.TrackNumber))
			if err != nil {
				logger.Errorf("getting parcel subscribers error: %s", err)
				continue
			}

			// publish notification for every subscriber
			for _, sub := range subs {
				// publish notification
				if err := n.deliveryProducer.Publish(&notificationpb.Delivery{
					UserTelegramId: int64(sub),
					TrackNumber:    event.TrackNumber,
					Checkpoint:     event.Checkpoint,
				}); err != nil {
					logger.Errorf("publishing notification error: %s", err)
				}
			}
		case <-ctx.Done():
			logger.Info("handler ended")
			return
		}
	}
}
