package state

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
)

type Register struct {
	Manager  model.Manager
	FillStep RegisterFillStep
}

type RegisterFillStep uint

const (
	RegisterFillStepEmpty RegisterFillStep = iota
	RegisterFillStepFullName
	RegisterFillStepEmail
	RegisterFillStepPhoneNumber
	RegisterFillStepCompany
	RegisterFillStepReady
)

func SetRegister(ss session.Session, tID model.TelegramID) {
	ss.SetState(Register{
		Manager: model.Manager{
			TelegramID: tID,
		},
	})
}

func (r *Register) Ended() bool {
	return r.FillStep == RegisterFillStepReady
}
