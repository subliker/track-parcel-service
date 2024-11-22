package bot

import (
	"context"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/session/state"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleOnText() tele.HandlerFunc {
	const handlerName = "on text"
	const errMsg = "handle " + handlerName + " error(session store): %s"
	logger := b.logger.WithFields("handler", handlerName)

	return func(ctx tele.Context) error {
		tID := model.TelegramID(ctx.Sender().ID)
		logger := logger.WithFields("user_id", tID)

		// getting state
		ss, err := b.sessionStore.Get(tID)

		// ignore text of session is not exist
		if err == session.ErrSessionIsNotExist {
			return nil
		}
		if err != nil {
			err := fmt.Errorf(errMsg, err)
			logger.Error(err)
			return err
		}

		switch st := ss.State().(type) {
		case state.MakeParcel:
			if err := b.fillParcel(ctx, &st); err != nil {
				err := fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			if st.FillStep == state.MakeParcelFillStepReady {
				ctx.Send("Посылка готова")
				logger.Info(st.Parcel)
				ss.ClearState()
				break
			}
			ss.SetState(st)
		case state.Register:
			if err := b.fillRegister(ctx, &st); err != nil {
				err := fmt.Errorf(errMsg, err)
				logger.Error(err)
				return err
			}
			if st.FillStep == state.RegisterFillStepReady {
				b.managerClient.Register(context.Background(), &managerpb.RegisterRequest{
					ManagerTelegramId:  int64(st.Manager.TelegramID),
					ManagerFullName:    st.Manager.FullName,
					ManagerEmail:       st.Manager.Email,
					ManagerPhoneNumber: st.Manager.PhoneNumber,
					ManagerCompany:     st.Manager.Company,
				})
				ctx.Send("Регистрация готова")
				logger.Info(st.Manager)
				ss.ClearState()
				break
			}
			ss.SetState(st)
		}

		return nil
	}
}
