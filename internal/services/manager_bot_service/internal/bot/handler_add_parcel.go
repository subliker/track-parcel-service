package bot

import (
	"fmt"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
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

		ctx.Send(b.bundle.States().MakeParcel().Name())
		return nil
	}
}

// TODO ADD LANG n VALIDATION n TIME PARSING
func (b *bot) fillParcel(ctx tele.Context, st *state.MakeParcel) error {
	// set state handler
	ctx.Set("state_handler", "fill parcel")

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

	return nil
}
