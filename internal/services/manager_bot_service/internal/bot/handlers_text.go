package bot

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

var dontSpecifyKetboard *tele.ReplyMarkup
var btnDontSpecify tele.Btn

func (b *bot) handleOnText() tele.HandlerFunc {
	dontSpecifyKetboard = b.client.NewMarkup()

	btnDontSpecify = dontSpecifyKetboard.Data(b.bundle.Common().Markup().BtnDontSpecify(), "dont-specify")

	dontSpecifyKetboard.Inline(dontSpecifyKetboard.Row(btnDontSpecify))
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "on text")

		tID := model.TelegramID(ctx.Sender().ID)

		// getting state
		ss, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("getting session error: %s", err)
		}
		switch st := ss.State().(type) {
		case state.MakeParcel:
			if err := b.fillParcel(ctx, &st); err != nil {
				return err
			}
			if st.Ended() {
				ss.ClearState()
				b.handleMenu()(ctx)
				break
			} else {
				ss.SetState(st)
			}
		case state.Register:
			if err := b.fillRegister(ctx, &st); err != nil {
				return err
			}
			if st.Ended() {
				ss.ClearState()
				b.handleMenu()(ctx)
				break
			} else {
				ss.SetState(st)
			}
		default:
			ctx.Send("Некорректный ввод")
		}

		return nil
	}
}

func (b *bot) handleDontSpecify() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		ctx.Set("dont-specify", true)
		ctx.Respond()
		return b.handleOnText()(ctx)
	}
}
