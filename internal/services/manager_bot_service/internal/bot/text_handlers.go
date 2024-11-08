package bot

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/models/telegram"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleOnText() tele.HandlerFunc {
	const handlerName = "on text"
	const errMsg = "handle " + handlerName + " error(session store): %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := telegram.ID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// getting state
		ss, err := b.sessionStore.Get(tID)

		// ignore text of session is not exist
		if err == session.ErrSessionIsNotExist {
			return nil
		}
		if err != nil {
			err := fmt.Errorf(errMsg, err)
			logger.Error(err)
			return err
		}

		switch st := ss.State().(type) {
		case state.MakeParcel:
			st, err := b.fillParcel(ctx, st)
			if err != nil {
				err := fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			if st.FillStep == state.MakeParcelFillStepReady {
				ctx.Send("Посылка готова")
				logger.Info(st.Parcel)
				// add parcel in
				ss.ClearState()
				break
			}
			ss.SetState(st)
		}

		return nil
	}
}
