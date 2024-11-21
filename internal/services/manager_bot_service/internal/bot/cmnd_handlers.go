package bot

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleStart() tele.HandlerFunc {
	const handlerName = "start"
	const errMsg = "handle start command error: %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// create user session if not exist
		if err := b.sessionStore.Add(model.TelegramID(ctx.Sender().ID)); err != nil {
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
	const errMsg = "handle add parcel command error: %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
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

func (b *bot) handleRegister() tele.HandlerFunc {
	const handlerName = "register"
	const errMsg = "handle register error: %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// register manager if not registered
		err := b.managerClient.Register(context.Background(), &managerpb.RegisterRequest{
			ManagerTelegramId: ctx.Sender().ID,
			ManagerFullName:   "Shcherbachev Andrey Nikolaevich",
			ManagerEmail:      "subliker0@gmail.com",
		})
		if errors.Is(err, manager.ErrInternal) {
			logger.Errorf(errMsg, err)
			return ctx.Send("internal error")
		}

		return nil
	}
}
