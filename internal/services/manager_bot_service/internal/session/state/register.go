package state

import "github.com/subliker/track-parcel-service/internal/pkg/model"

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
