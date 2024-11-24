package bot

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pmpb"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	"google.golang.org/protobuf/types/known/timestamppb"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleOnText() tele.HandlerFunc {

	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "on text")

		tID := model.TelegramID(ctx.Sender().ID)

		// getting state
		ss, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("getting session error: %s", err)
		}
		switch st := ss.State().(type) {
		case state.MakeParcel:
			if err := b.fillParcel(ctx, &st); err != nil {
				return err
			}
			if st.Ended() {
				p := st.Parcel
				res, err := b.parcelsManagerClient.AddParcel(context.Background(), &pmpb.AddParcelRequest{
					ParcelName:           p.Name,
					ManagerTelegramId:    int64(p.ManagerID),
					ParcelRecipient:      p.Recipient,
					ParcelArrivalAddress: p.ArrivalAddress,
					ParcelForecastDate:   timestamppb.New(p.ForecastDate),
					ParcelDescription:    p.Description,
				})
				if err != nil {
					ctx.Send("register ended with internal error")
					return err
				}
				ctx.Send(b.bundle.States().MakeParcel().Ready(res.TrackNumber))
				ss.ClearState()
				break
			} else {
				ss.SetState(st)
			}
		case state.Register:
			if err := b.fillRegister(ctx, &st); err != nil {
				return err
			}
			if st.Ended() {
				ss.ClearState()
				break
			} else {
				ss.SetState(st)
			}
		default:
			ctx.Send("Некорректный ввод")
		}

		return nil
	}
}
