package bot

import (
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) fillRegister(ctx tele.Context, st *state.Register) error {
	// set state handler
	ctx.Set("state_handler", "fill register")

	st.FillStep++

	fillBundle := b.bundle.States().Register()
	switch st.FillStep {
	case state.RegisterFillStepFullName:
		st.Manager.FullName = ctx.Text()
		ctx.Send(fillBundle.Email())
	case state.RegisterFillStepEmail:
		st.Manager.Email = ctx.Text()
		ctx.Send(fillBundle.PhoneNumber())
	case state.RegisterFillStepPhoneNumber:
		if ctx.Text() == "NO" {
			st.Manager.PhoneNumber = nil
		} else {
			t := ctx.Text()
			st.Manager.PhoneNumber = &t
		}
		ctx.Send(fillBundle.Company())
	case state.RegisterFillStepCompany:
		if ctx.Text() == "NO" {
			st.Manager.Company = nil
		} else {
			t := ctx.Text()
			st.Manager.Company = &t
		}
		st.FillStep = state.RegisterFillStepReady
	}

	return nil
}
