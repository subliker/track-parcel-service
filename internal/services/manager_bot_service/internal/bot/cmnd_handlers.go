package bot

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/bot/style"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleStart() tele.HandlerFunc {
	// logger := b.logger.WithFields("handler", "start")

	return func(ctx tele.Context) error {
		// tID := model.TelegramID(ctx.Sender().ID)
		// _ = logger.WithFields("user_id", tID)

		ctx.Send(b.bundle.OnStartMessage(ctx.Sender().FirstName), style.MenuKeyboard)
		return nil
	}
}

func (b *bot) handleAddParcel() tele.HandlerFunc {
	logger := b.logger.WithFields("handler", "add parcel")

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			logger.Errorf("get session error: %s", err)
			return ctx.Send("internal error")
		}

		// set make parcel state
		state.SetMakeParcel(session)

		ctx.Send(b.bundle.States().MakeParcel().Name())
		return nil
	}
}

func (b *bot) handleRegister() tele.HandlerFunc {
	logger := b.logger.WithFields("handler", "register")

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// check if authorized
		authorized, ok := ctx.Get("authorized").(bool)
		if !ok {
			logger.Error("auth error: authorized is nil")
			return ctx.Send("internal error")
		}

		if authorized {
			return ctx.Send("manager have been already registered")
		}

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			logger.Errorf("get session error: %s", err)
			return ctx.Send("internal error")
		}

		// set register state
		state.SetRegister(session, tID)

		return nil
	}
}
