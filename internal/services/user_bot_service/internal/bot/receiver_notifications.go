package bot

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) receiveNotification() {
	logger := b.logger.WithFields("handler", "notifications receiver")

	logger.Info("notifications receiver started")
	for not := range b.deliveryConsumer.Listen() {
		// TODO fix time format
		_, err := b.client.Send(&tele.User{ID: not.UserTelegramId}, b.bundle.Notification().Main(
			not.TrackNumber, not.Checkpoint.Time.AsTime().Format(model.ForecastDateLayout),
			not.Checkpoint.Place, not.Checkpoint.Description, not.Checkpoint.ParcelStatus.String(),
		))
		if err != nil {
			logger.Error(err)
		}
	}
	logger.Info("notifications receiver stopped")
}
