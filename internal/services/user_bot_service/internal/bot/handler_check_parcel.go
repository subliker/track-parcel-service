package bot

import (
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

var showParcelBtnRefresh tele.Btn

func (b *bot) handleCheckParcel() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "check parcel")

		tID := model.TelegramID(ctx.Sender().ID)

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("get session error: %s", err)
		}

		// set check parcel state
		state.SetCheckParcel(session)
		ctx.Send("enter track number")

		return nil
	}
}

func (b *bot) handleShowParcelRefreshBtn() tele.HandlerFunc {
	showParcelBtnRefresh = b.client.NewMarkup().Data("refresh", "show-parcel-refresh")
	return func(ctx tele.Context) error {
		return nil
	}
}
func (b *bot) handleShowParcel(ctx tele.Context) error {
	// set handler name
	ctx.Set("handler", "show parcel")
	// getting parcel track number
	trackNum, ok := ctx.Get("parcel-track-number").(model.TrackNumber)
	if !ok {
		return errors.New("error getting parcel track number from context")
	}
	// getting parcel struct
	p, ok := ctx.Get("parcel").(model.Parcel)
	if !ok {
		return errors.New("error getting parcel from context")
	}
	// making markup
	mu := b.client.NewMarkup()
	btn := showParcelBtnRefresh
	btn.Data = string(trackNum)

	mu.Inline(mu.Row(btn))
	// showing parcel
	ctx.Send(b.bundle.CheckParcel().Main(
		p.Name, p.Recipient, p.ArrivalAddress,
		p.ForecastDate.String(),
		p.Description, string(p.Status),
	), mu)
	return nil
}

func (b *bot) onCheckParcelState(
	ctx tele.Context, ss session.Session,
	st state.CheckParcel,
) error {
	// make fill iteration
	ended, err := st.Next(ctx.Text())
	if err != nil {
		return nil
	}

	// send
	if ended {
		err := st.Ready(
			b.parcelsUserClient,
			func(text string) {
				ctx.Send(text)
			},
			b.bundle,
		)
		if err != nil {
			ctx.Send("internal error")
		}
		ss.ClearState()
		b.handleMenu()(ctx)
		return err
	} else {
		ss.SetState(st)
	}
	return nil
}
