package state

import models "github.com/subliker/track-parcel-service/internal/pkg/models/parcel"

type MakeParcel struct {
	Parcel   models.Parcel
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
