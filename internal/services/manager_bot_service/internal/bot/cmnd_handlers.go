package bot

import (
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot/style"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleStart() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "start")

		ctx.Send(b.bundle.OnStartMessage(ctx.Sender().FirstName), style.MenuKeyboard)
		return nil
	}
}

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

		// set make parcel state
		state.SetMakeParcel(session)

		ctx.Send(b.bundle.States().MakeParcel().Name())
		return nil
	}
}

func (b *bot) handleRegister() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "register")

		tID := model.TelegramID(ctx.Sender().ID)

		// check if authorized
		authorized, ok := ctx.Get("authorized").(bool)
		if !ok {
			ctx.Send("internal error")
			return errors.New("auth error: authorized is nil")
		}

		if authorized {
			ctx.Send("manager have been already registered")
			return nil
		}

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("get session error: %s", err)
		}

		// set register state
		state.SetRegister(session, tID)
		ctx.Send(b.bundle.States().Register().FullName())

		return nil
	}
}
