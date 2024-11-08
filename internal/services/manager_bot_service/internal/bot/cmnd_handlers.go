package bot

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleStart() tele.HandlerFunc {
	const handlerName = "start"
	const errMsg = "handle " + handlerName + " error(session store): %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := telegram.ID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// create user session if not exist
		if err := b.sessionStore.Add(telegram.ID(ctx.Sender().ID)); err != nil {
			if err != session.ErrSessionIsAlreadyExist {
				err := fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
		}

		ctx.Send(b.bundle.OnStartMessage(ctx.Sender().FirstName))
		return nil
	}
}

func (b *bot) handleAddParcel() tele.HandlerFunc {
	const handlerName = "add parcel"
	const errMsg = "handle " + handlerName + " error(session store): %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := telegram.ID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// create user session if not exist
		if err := b.sessionStore.Add(tID); err != nil {
			if err != session.ErrSessionIsAlreadyExist {
				err := fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
		}

		// set make parcel state
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			err := fmt.Errorf(errMsg, err)
			logger.Error(err)
			return err
		}

		session.SetState(state.MakeParcel{})

		ctx.Send(b.bundle.States().MakeParcel().Name())
		return nil
	}
}
