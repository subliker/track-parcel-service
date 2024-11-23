package bot

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
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
				ctx.Send(b.bundle.States().MakeParcel().Ready())
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
				err := b.managerClient.Register(context.Background(), &managerpb.RegisterRequest{
					ManagerTelegramId:  int64(st.Manager.TelegramID),
					ManagerFullName:    st.Manager.FullName,
					ManagerEmail:       st.Manager.Email,
					ManagerPhoneNumber: st.Manager.PhoneNumber,
					ManagerCompany:     st.Manager.Company,
				})
				if err == manager.ErrManagerIsAlreadyExist {
					ctx.Send("you have been already registered")
					return nil
				}
				if err != nil {
					ctx.Send("register ended with internal error")
					return nil
				}
				ctx.Send(b.bundle.States().Register().Ready())
				ss.ClearState()
				break
			} else {
				ss.SetState(st)
			}
		}

		return nil
	}
}
