package bot

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	models "github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleStart() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		if err := b.sessionStore.Add(models.TelegramID(ctx.Sender().ID)); err != nil {
			if err != session.ErrSessionIsAlreadyExist {
				err := fmt.Errorf("handle start error(session store): %s", err)
				logger.Zap.Error(err)
				return err
			}
		}

		ctx.Send(b.bundle.OnStartMessage(ctx.Sender().FirstName))
		return nil
	}
}

func (b *bot) handleAddParcel() tele.HandlerFunc {
	return func(ctx tele.Context) error {

		ctx.Send(b.bundle.OnSetParcel(""))
		return nil
	}
}
