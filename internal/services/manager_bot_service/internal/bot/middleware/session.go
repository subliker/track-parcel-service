package middleware

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	tele "gopkg.in/telebot.v4"
)

func Session(logger logger.Logger, sessionStore session.Store) tele.MiddlewareFunc {
	logger = logger.WithFields("middleware", "session")
	const errMsg = "ensure session error: %s"
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if err := sessionStore.Ensure(model.TelegramID(ctx.Sender().ID)); err != nil {
				ctx.Send("internal error")
				err = fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			return next(ctx)
		}
	}
}
