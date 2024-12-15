package bot

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
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

		ctx.Send(b.bundle.MyApi().Main(string(apiToken)))
		return nil
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
