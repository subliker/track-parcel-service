package middleware

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/lang"
	tele "gopkg.in/telebot.v4"
)

func Auth(logger logger.Logger, managerClient manager.Client, bundle lang.Messages) tele.MiddlewareFunc {
	logger = logger.WithFields("middleware", "auth")
	const errMsg = "manager auth error: %s"
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			// try to auth by telegram id
			err := managerClient.Auth(context.Background(), &managerpb.AuthRequest{
				ManagerTelegramId: int64(ctx.Sender().ID),
			})
			// if error exists and isn't ErrManagerNotFound throw error
			if err != nil && err != manager.ErrManagerNotFound {
				ctx.Send(bundle.Common().Errors().Internal())
				err = fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			// authorized if err in't manager.ErrManagerNotFound
			ctx.Set("authorized", err != manager.ErrManagerNotFound)
			return next(ctx)
		}
	}
}

func Authorized(logger logger.Logger, bundle lang.Messages) tele.MiddlewareFunc {
	logger = logger.WithFields("middleware", "authorized")
	const errMsg = "manager authorized error: %s"
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
