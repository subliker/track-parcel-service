package bot

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleAddParcel() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "add parcel")

		tID := model.TelegramID(ctx.Sender().ID)

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("get session error: %s", err)
		}

		// set add parcel state
		state.SetAddParcel(session)

		ctx.Send(b.bundle.AddParcel().Points().Name())
		return nil
	}
}

func (b *bot) onAddParcelState(
	ctx tele.Context, ss session.Session,
	st state.AddParcel,
) error {
	// make fill iteration
	ended, err := st.Next(
		ctx.Text(),
		func(text string) {
			ctx.Send(text)
		},
		b.bundle,
	)
	if err != nil {
		return err
	}
	// send
	if ended {
		err := st.Ready(
			b.parcelsManagerClient,
			func(text string) {
				ctx.Send(text)
			},
			b.bundle,
		)
		if err != nil {
			return err
		}
		ss.ClearState()
		b.handleMenu()(ctx)
		return err
	} else {
		ss.SetState(st)
	}
	return nil
}
