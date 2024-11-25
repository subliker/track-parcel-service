package bot

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

var checkParcelBtnRefresh tele.Btn

func (b *bot) handleCheckParcel() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "check parcel")

		tID := model.TelegramID(ctx.Sender().ID)

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("get session error: %s", err)
		}

		// set check parcel state
		state.SetRegister(session, tID)
		ctx.Send(b.bundle.Register().Points().FullName())

		return nil
	}
}

func (b *bot) fillCheckParcel(ctx tele.Context, st *state.CheckParcel) error {
	// set state handler
	ctx.Set("state_handler", "fill register")

	st.TrackNum = ctx.Text()
	return nil
}

func (b *bot) sendCheckParcel(ctx tele.Context, trackNum string) error {
	return nil
}
