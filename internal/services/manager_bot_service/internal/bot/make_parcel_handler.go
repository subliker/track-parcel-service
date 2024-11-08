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
	switch st.FillStep {
	case state.MakeParcelFillStepName:
		st.Parcel.Name = ctx.Text()
		ctx.Send("Введите имя получателя")
	case state.MakeParcelFillRecipient:
		st.Parcel.Recipient = ctx.Text()
		ctx.Send("Введите адрес получения")
	case state.MakeParcelFillArrivalAddress:
		st.Parcel.ArrivalAddress = ctx.Text()
		ctx.Send("Введите предположительную дату доставки")
	case state.MakeParcelFillForecastDate:
		fd, err := time.Parse(time.RFC3339, ctx.Text())
		fmt.Print(ctx.Text(), err)
		if err != nil {
			// undo step
			st.FillStep--
			ctx.Reply("Введено некорректное время")
			break
		}
		st.Parcel.ForecastDate = fd
		ctx.Send("Введите описание посылки")
	case state.MakeParcelFillDescription:
		st.Parcel.Description = ctx.Text()
		st.FillStep = state.MakeParcelFillStepReady
	}

	return st, nil
}
