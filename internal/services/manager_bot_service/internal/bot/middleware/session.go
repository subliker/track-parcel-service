package middleware

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	tele "gopkg.in/telebot.v4"
)

func Session(sessionStore session.Store) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if err := sessionStore.Ensure(model.TelegramID(ctx.Sender().ID)); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}
