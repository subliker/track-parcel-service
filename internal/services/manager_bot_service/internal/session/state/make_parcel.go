package state

import (
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
)

const StateTypeMakeParcel = "make parcel"

type MakeParcel struct {
	Parcel   model.Parcel
	FillStep MakeParcelFillStep
}

type MakeParcelFillStep uint

const (
	MakeParcelFillStepEmpty MakeParcelFillStep = iota
	MakeParcelFillStepName
	MakeParcelFillRecipient
	MakeParcelFillArrivalAddress
	MakeParcelFillForecastDate
	MakeParcelFillDescription
	MakeParcelFillStepReady
)

func SetMakeParcel(ss session.Session) {
	ss.SetState(&MakeParcel{})
}

func (mp *MakeParcel) Ended() bool {
	return mp.FillStep == MakeParcelFillStepReady
}
