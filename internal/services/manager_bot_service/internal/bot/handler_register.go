package bot

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleRegister() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "register")

		tID := model.TelegramID(ctx.Sender().ID)

		// if on callback
		ctx.Respond()

		// check if authorized
		authorized, ok := ctx.Get("authorized").(bool)
		if !ok {
			ctx.Send("internal error")
			return errors.New("auth error: authorized is nil")
		}

		if authorized {
			ctx.Send("manager have been already registered")
			return nil
		}

		// get session
		session, err := b.sessionStore.Get(tID)
		if err != nil {
			ctx.Send("internal error")
			return fmt.Errorf("get session error: %s", err)
		}

		// set register state
		state.SetRegister(session, tID)
		ctx.Send(b.bundle.Register().Points().FullName())

		return nil
	}
}

func (b *bot) fillRegister(ctx tele.Context, st *state.Register) error {
	// set state handler
	ctx.Set("state_handler", "fill register")

	st.FillStep++

	fillBundle := b.bundle.Register().Points()
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
		fallthrough
	case state.RegisterFillStepReady:
		err := b.sendRegister(ctx, st.Manager)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *bot) sendRegister(ctx tele.Context, m model.Manager) error {
	err := b.managerClient.Register(context.Background(), &managerpb.RegisterRequest{
		ManagerTelegramId:  int64(m.TelegramID),
		ManagerFullName:    m.FullName,
		ManagerEmail:       m.Email,
		ManagerPhoneNumber: m.PhoneNumber,
		ManagerCompany:     m.Company,
	})
	if errors.Is(err, manager.ErrManagerIsAlreadyExist) {
		ctx.Send("you have been already registered")
		return err
	}
	if err != nil {
		ctx.Send("register ended with internal error")
		return err
	}

	// optional
	var phoneNumber, company string
	if m.PhoneNumber != nil {
		phoneNumber = *m.PhoneNumber
	}
	if m.Company != nil {
		company = *m.Company
	}

	ctx.Send(b.bundle.Register().Points().Ready(m.FullName, m.Email, phoneNumber, company))
	return nil
}
