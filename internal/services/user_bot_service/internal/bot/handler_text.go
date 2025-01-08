package bot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

var btnNotSpecify tele.Btn

func (b *bot) notSpecifyKeyboard(data string) *tele.ReplyMarkup {
	k := b.client.NewMarkup()
	btn := btnNotSpecify
	btn.Data = data
	k.Inline(k.Row(btn))
	return k
}

func (b *bot) handleOnText() tele.HandlerFunc {
	btnNotSpecify = (&tele.ReplyMarkup{}).Data(b.bundle.Common().Markup().BtnDontSpecify(), "not-specify", "0")
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "on text")

		tID := model.TelegramID(ctx.Sender().ID)

		// get not specify
		notSpecifyField, _ := ctx.Get("not-specify-field").(uint)

		// getting state
		ss, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send(b.bundle.Common().Errors().Internal())
			return fmt.Errorf("getting session error: %s", err)
		}
		switch st := ss.State().(type) {
		case state.Register:
			return b.onRegisterState(ctx, ss, st, notSpecifyField)
		case state.CheckParcel:
			if notSpecifyField > 0 {
				break
			}
			return b.onCheckParcelState(ctx, ss, st)
		default:
			ctx.Send(b.bundle.Common().Errors().IncorrectInput())
		}

		return nil
	}
}

func (b *bot) handleDontSpecify() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		data := ctx.Callback().Data
		field, err := strconv.Atoi(data)
		if err != nil {
			return errors.New("incorrect callback data")
		}

		ctx.Set("not-specify-field", uint(field))
		ctx.Respond()
		return b.handleOnText()(ctx)
	}
}
