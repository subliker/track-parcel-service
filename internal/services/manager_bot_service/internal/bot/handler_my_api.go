package bot

import (
	"context"
	"fmt"
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleMyApi() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "my api")

		tID := model.TelegramID(ctx.Sender().ID)

		// get token
		apiToken, err := b.getApiToken(tID)
		if err != nil {
			return fmt.Errorf("api token request ended with: %s", err)
		}

		mk := b.client.NewMarkup()
		mk.Inline(mk.Row(mk.URL(b.bundle.MyApi().Btns().Docs(), b.apiTarget+"/swagger")))
		return ctx.Send(b.bundle.MyApi().Main(string(apiToken), b.apiTarget), mk)
	}
}

func (b *bot) getApiToken(managerTID model.TelegramID) (model.ManagerApiToken, error) {
	// request api token
	res, err := b.managerClient.GetApiToken(context.Background(), &managerpb.GetApiTokenRequest{
		ManagerTelegramId: int64(managerTID),
	})
	if err != nil {
		return "", err
	}
	return model.ManagerApiToken(res.ManagerApiToken), nil
}
