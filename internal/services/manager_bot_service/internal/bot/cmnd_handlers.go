package bot

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleStart() tele.HandlerFunc {
	logger := b.logger.WithFields("handler", "start")

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// ensure session exists
		if err := b.sessionStore.Ensure(tID); err != nil {
			const errMsg = "error ensuring user session exists: %s"
			err := fmt.Errorf(errMsg, err)
			logger.Error(err)
			return err
		}

		ctx.Send(b.bundle.OnStartMessage(ctx.Sender().FirstName))
		return nil
	}
}

func (b *bot) handleAddParcel() tele.HandlerFunc {
	logger := b.logger.WithFields("handler", "add parcel")

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// ensure get session
		session, err := b.sessionStore.EnsureGet(tID)
		if err != nil {
			logger.Errorf("ensure get session error: %s", err)
			return ctx.Send("internal error")
		}

		session.SetState(state.MakeParcel{})

		ctx.Send(b.bundle.States().MakeParcel().Name())
		return nil
	}
}

func (b *bot) handleRegister() tele.HandlerFunc {
	logger := b.logger.WithFields("handler", "register")

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// try to auth by telegram id
		err := b.managerClient.Auth(context.Background(), &managerpb.AuthRequest{
			ManagerTelegramId: int64(tID),
		})
		if err == nil {
			return ctx.Send("manager have been already registered")
		}
		if err != manager.ErrManagerNotFound {
			logger.Errorf("manager auth error: %s", err)
			return ctx.Send("internal error")
		}

		// ensure get session
		session, err := b.sessionStore.EnsureGet(tID)
		if err != nil {
			logger.Errorf("ensure get session error: %s", err)
			return ctx.Send("internal error")
		}

		session.SetState(state.Register{
			Manager: model.Manager{
				TelegramID: tID,
			},
		})

		return nil
	}
}
