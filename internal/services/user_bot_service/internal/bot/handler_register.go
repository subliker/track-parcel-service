package bot

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/userpb"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/session/state"
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
			ctx.Send("user have been already registered")
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

	notSpecify, _ := ctx.Get("dont-specify").(bool)

	st.FillStep++

	fillBundle := b.bundle.Register().Points()
	switch st.FillStep {
	case state.RegisterFillStepFullName:
		st.User.FullName = ctx.Text()
		ctx.Send(fillBundle.Email())
	case state.RegisterFillStepEmail:
		st.User.Email = ctx.Text()

		ctx.Send(fillBundle.PhoneNumber(), dontSpecifyKetboard)
	case state.RegisterFillStepPhoneNumber:
		if notSpecify {
			st.User.PhoneNumber = nil
		} else {
			t := ctx.Text()
			st.User.PhoneNumber = &t
		}
		st.FillStep++
		fallthrough
	case state.RegisterFillStepReady:
		err := b.sendRegister(ctx, st.User)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *bot) sendRegister(ctx tele.Context, u model.User) error {
	err := b.userClient.Register(context.Background(), &userpb.RegisterRequest{
		UserTelegramId:  int64(u.TelegramID),
		UserFullName:    u.FullName,
		UserEmail:       u.Email,
		UserPhoneNumber: u.PhoneNumber,
	})
	if errors.Is(err, user.ErrUserIsAlreadyExist) {
		ctx.Send("you have been already registered")
		return err
	}
	if err != nil {
		ctx.Send("register ended with internal error")
		return err
	}

	// optional
	var phoneNumber string
	if u.PhoneNumber != nil {
		phoneNumber = *u.PhoneNumber
	}

	ctx.Send(b.bundle.Register().Points().Ready(u.FullName, u.Email, phoneNumber))
	return nil
}
