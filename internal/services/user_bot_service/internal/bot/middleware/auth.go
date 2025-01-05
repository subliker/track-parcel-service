package middleware

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/userpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/lang"
	tele "gopkg.in/telebot.v4"
)

func Auth(logger logger.Logger, userClient user.Client, bundle lang.Messages) tele.MiddlewareFunc {
	logger = logger.WithFields("middleware", "auth")
	const errMsg = "user auth error: %s"
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			// try to auth by telegram id
			err := userClient.Auth(context.Background(), &userpb.AuthRequest{
				UserTelegramId: int64(ctx.Sender().ID),
			})
			// if error exists and isn't ErrUserNotFound throw error
			if err != nil && err != user.ErrUserNotFound {
				ctx.Send(bundle.Common().Errors().Internal())
				err = fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			// authorized if err isn't user.ErrUserNotFound
			ctx.Set("authorized", err != user.ErrUserNotFound)
			return next(ctx)
		}
	}
}

func Authorized(logger logger.Logger, bundle lang.Messages) tele.MiddlewareFunc {
	logger = logger.WithFields("middleware", "authorized")
	const errMsg = "user authorized error: %s"
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			authorized, ok := ctx.Get("authorized").(bool)
			if !ok {
				ctx.Send(bundle.Common().Errors().Internal())
				err := fmt.Errorf(errMsg, "auth middleware dropped")
				logger.Error(err)
				return err
			}
			if !authorized {
				ctx.Send(bundle.Common().Errors().NotAuthorized())
				return nil
			}
			return next(ctx)
		}
	}
}
