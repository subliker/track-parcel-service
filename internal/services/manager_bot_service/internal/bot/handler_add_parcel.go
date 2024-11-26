package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pmpb"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	"google.golang.org/protobuf/types/known/timestamppb"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleAddParcel() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "add parcel")

		tID := model.TelegramID(ctx.Sender().ID)

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("get session error: %s", err)
		}

		// set make parcel state
		state.SetMakeParcel(session)

		ctx.Send(b.bundle.AddParcel().Points().Name())
		return nil
	}
}

// TODO ADD LANG n VALIDATION n TIME PARSING
func (b *bot) fillParcel(ctx tele.Context, st *state.MakeParcel) error {
	// set state handler
	ctx.Set("state_handler", "fill parcel")

	st.FillStep++

	fillBundle := b.bundle.AddParcel().Points()
	switch st.FillStep {
	case state.MakeParcelFillStepName:
		st.Parcel.Name = ctx.Text()
		ctx.Send(fillBundle.Recipient())
	case state.MakeParcelFillRecipient:
		st.Parcel.Recipient = ctx.Text()
		ctx.Send(fillBundle.ArrivalAddress())
	case state.MakeParcelFillArrivalAddress:
		st.Parcel.ArrivalAddress = ctx.Text()
		ctx.Send(fillBundle.ForecastDate())
	case state.MakeParcelFillForecastDate:
		fd, err := time.Parse(time.RFC3339, ctx.Text())
		fmt.Print(ctx.Text(), err)
		if err != nil {
			// undo step
			st.FillStep--
			// ctx.Reply(fillBundle.ForecastDateIncorrectTime())
			break
		}
		st.Parcel.ForecastDate = fd
		b.logger.Error(ctx.Send(fillBundle.Description()))
	case state.MakeParcelFillDescription:
		st.Parcel.Description = ctx.Text()

		st.FillStep++
	}

	return nil
}

func (b *bot) sendParcel(ctx tele.Context, p model.Parcel) error {
	res, err := b.parcelsManagerClient.AddParcel(context.Background(), &pmpb.AddParcelRequest{
		ParcelName:           p.Name,
		ManagerTelegramId:    int64(p.ManagerID),
		ParcelRecipient:      p.Recipient,
		ParcelArrivalAddress: p.ArrivalAddress,
		ParcelForecastDate:   timestamppb.New(p.ForecastDate),
		ParcelDescription:    p.Description,
	})
	if err != nil {
		ctx.Send("add parcel ended with internal error")
		return err
	}

	ctx.Send(b.bundle.AddParcel().Points().Ready(p.Name, p.Recipient, p.ArrivalAddress, p.ForecastDate.String(), p.Description, res.TrackNumber))
	return nil
}
