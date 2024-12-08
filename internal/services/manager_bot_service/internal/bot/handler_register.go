package bot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleRegister() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "register")

		tID := model.TelegramID(ctx.Sender().ID)

		// if on callback
		ctx.Respond()

		// check if authorized
		authorized, ok := ctx.Get("authorized").(bool)
		if !ok {
			ctx.Send(b.bundle.Common().Errors().Internal())
			return errors.New("auth error: authorized is nil")
		}

		if authorized {
			ctx.Send("manager have been already registered")
			return nil
		}

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send(b.bundle.Common().Errors().Internal())
			return fmt.Errorf("get session error: %s", err)
		}

		// set register state
		state.SetRegister(session, tID)

		ctx.Send(b.bundle.Register().Points().FullName())
		return nil
	}
}

func (b *bot) onRegisterState(
	ctx tele.Context, ss session.Session,
	st state.Register, notSpecifyField uint,
) error {
	// make fill iteration
	ended, err := st.Next(
		ctx.Text(),
		func(text string, optionalField state.RegisterFillStep) {
			if optionalField > 0 {
				ctx.Send(text, b.notSpecifyKeyboard(strconv.Itoa(int(optionalField))))
				return
			}
			ctx.Send(text)
		},
		b.bundle,
		state.RegisterFillStep(notSpecifyField),
	)
	// ignore incorrect not specify
	if err != nil && err != session.ErrIncorrectNotSpecify {
		return err
	}
	// send
	if ended {
		err := st.Ready(
			b.managerClient,
			func(text string) {
				ctx.Send(text)
			},
			b.bundle,
		)
		if err != nil {
			ctx.Send(b.bundle.Common().Errors().Internal())
		}
		ss.ClearState()
		b.handleMenu()(ctx)
		return err
	} else {
		ss.SetState(st)
	}
	return nil
}
