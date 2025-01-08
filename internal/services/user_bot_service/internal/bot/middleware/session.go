package middleware

import (
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/lang"
	tele "gopkg.in/telebot.v4"
)

func Session(logger logger.Logger, sessionStore session.Store, bundle lang.Messages) tele.MiddlewareFunc {
	logger = logger.WithFields("middleware", "session")
	const errMsg = "ensure session error: %s"
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			// ensure user session
			if err := sessionStore.Ensure(model.TelegramID(ctx.Sender().ID)); err != nil {
				ctx.Send(bundle.Common().Errors().Internal())
				err = fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			return next(ctx)
		}
	}
}
