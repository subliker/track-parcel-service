package bot

import (
	"fmt"
	"time"

	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

// TODO ADD LANG n VALIDATION n TIME PARSING
func (b *bot) fillParcel(ctx tele.Context, st state.MakeParcel) (state.MakeParcel, error) {
	st.FillStep++

	fillBundle := b.bundle.States().MakeParcel()
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
			ctx.Reply(fillBundle.ForecastDateIncorrectTime())
			break
		}
		st.Parcel.ForecastDate = fd
		ctx.Send(fillBundle.Description())
	case state.MakeParcelFillDescription:
		st.Parcel.Description = ctx.Text()
		st.FillStep = state.MakeParcelFillStepReady
	}

	return st, nil
}
