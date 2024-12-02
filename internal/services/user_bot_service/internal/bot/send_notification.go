package bot

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) sendNotification(c model.Checkpoint, userTID model.TelegramID) {
	b.client.Send(&tele.User{
		ID: int64(userTID),
	}, c.Description)
}
