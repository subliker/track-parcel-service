package state

import (
	"context"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pm"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/parcelpb"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/pmpb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/lang"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MakeParcel is state to request parcel data to create one
type AddParcel struct {
	Parcel   model.Parcel
	FillStep AddParcelFillStep
}

// AddParcelFillStep is enum to mention every fill step
type AddParcelFillStep uint

const (
	AddParcelFillStepEmpty AddParcelFillStep = iota
	AddParcelFillStepName
	AddParcelFillRecipient
	AddParcelFillArrivalAddress
	AddParcelFillForecastDate
	AddParcelFillDescription
	AddParcelFillStepReady
)

// SetAddParcel sets empty add parcel state in user session
func SetAddParcel(ss session.Session) {
	ss.SetState(AddParcel{})
}

// done returns true if all state data is filled up
func (ap *AddParcel) done() bool {
	return ap.FillStep == AddParcelFillStepReady
}

// Next starts one of iterations to fill state up and returns true if all state data is filled up.
func (ap *AddParcel) Next(
	text string,
	request func(text string),
	bundle lang.Messages,
) (bool, error) {
	// increment step
	ap.FillStep++

	fillBundle := bundle.AddParcel().Points()
	errBundle := bundle.Common().Errors()

	switch ap.FillStep {
	case AddParcelFillStepName:
		// check length
		if err := validator.V.Var(text, "min=3,max=100"); err != nil {
			// send error
			request(errBundle.Length(3, 100))
			// undo
			ap.FillStep--
			break
		}
		// fill
		ap.Parcel.Name = text
		// request next step
		request(fillBundle.Recipient())
	case AddParcelFillRecipient:
		// check length
		if err := validator.V.Var(text, "min=3,max=255"); err != nil {
			// send error
			request(errBundle.Length(3, 255))
			// undo
			ap.FillStep--
			break
		}
		// fill
		ap.Parcel.Recipient = text
		// request next step
		request(fillBundle.ArrivalAddress())
	case AddParcelFillArrivalAddress:
		// check length
		if err := validator.V.Var(text, "min=3,max=255"); err != nil {
			// send error
			request(errBundle.Length(3, 255))
			// undo
			ap.FillStep--
			break
		}
		// fill
		ap.Parcel.ArrivalAddress = text
		// request next step
		request(fillBundle.ForecastDate())
	case AddParcelFillForecastDate:
		// try to parse time
		fd, err := time.Parse(model.ForecastDateLayout, text)
		if err != nil {
			// send error
			request(errBundle.TimeFormat())
			// undo
			ap.FillStep--
			break
		}
		// fill
		ap.Parcel.ForecastDate = fd
		// request next step
		request(fillBundle.Description())
	case AddParcelFillDescription:
		// check length
		if err := validator.V.Var(text, "max=255"); err != nil {
			// send error
			request(errBundle.Length(0, 255))
			// undo
			ap.FillStep--
			break
		}
		// fill
		ap.Parcel.Description = text
		// jump to ready
		ap.FillStep++
	}
	return ap.done(), nil
}

// Ready completes all data and send request
func (ap *AddParcel) Ready(
	parcelsManagerClient pm.Client,
	send func(text string),
	bundle lang.Messages,
) error {
	// check if state done
	if !ap.done() {
		return session.ErrStateNotDone
	}

	// request add parcel
	p := ap.Parcel
	res, err := parcelsManagerClient.AddParcel(context.Background(), &pmpb.AddParcelRequest{
		Parcel: &parcelpb.Parcel{
			Name:              p.Name,
			ManagerTelegramId: int64(p.ManagerID),
			Recipient:         p.Recipient,
			ArrivalAddress:    p.ArrivalAddress,
			ForecastDate:      timestamppb.New(p.ForecastDate),
		},
	})
	if err != nil {
		send(bundle.Common().Errors().Internal())
		return err
	}

	send(bundle.AddParcel().Points().Ready(
		p.Name,
		p.Recipient,
		p.ArrivalAddress,
		p.ForecastDate.Format(model.ForecastDateLayout),
		p.Description,
		res.TrackNumber,
	))
	return nil
}
