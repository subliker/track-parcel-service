package middleware

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	tele "gopkg.in/telebot.v4"
)

func Auth(logger logger.Logger, managerClient manager.Client) tele.MiddlewareFunc {
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
				ctx.Send("internal error")
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
