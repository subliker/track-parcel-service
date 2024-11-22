package state

import "github.com/subliker/track-parcel-service/internal/pkg/model"

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
