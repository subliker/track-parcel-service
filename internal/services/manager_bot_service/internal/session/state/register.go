package state

import (
	"context"
	"errors"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/lang"
)

// Register is state to request register manager from manager service
type Register struct {
	Manager  model.Manager
	FillStep RegisterFillStep
}

// RegisterFillStep is enum to mention every fill step
type RegisterFillStep uint

const (
	RegisterFillStepEmpty RegisterFillStep = iota
	RegisterFillStepFullName
	RegisterFillStepEmail
	RegisterFillStepPhoneNumber
	RegisterFillStepCompany
	RegisterFillStepReady
)

// SetRegister sets empty register state in user session
func SetRegister(ss session.Session, tID model.TelegramID) {
	ss.SetState(Register{
		Manager: model.Manager{
			TelegramID: tID,
		},
	})
}

// done returns true if all state data is filled up
func (r *Register) done() bool {
	return r.FillStep == RegisterFillStepReady
}

// Next starts one of iterations to fill state up and returns true if all state data is filled up.
// Ignores notSpecify if field is not optional
func (r *Register) Next(
	text string,
	request func(text string, optionalField RegisterFillStep),
	bundle lang.Messages,
	notSpecify RegisterFillStep,
) (bool, error) {
	// check not specify
	if notSpecify > RegisterFillStepEmpty && r.FillStep+1 != notSpecify {
		return false, session.ErrIncorrectNotSpecify
	}

	// increment step
	r.FillStep++

	skip := notSpecify > RegisterFillStepEmpty

	// lang bundle
	fillBundle := bundle.Register().Points()
	errBundle := bundle.Common().Errors()

	switch r.FillStep {
	case RegisterFillStepFullName:
		// check length
		if err := validator.V.Var(text, "min=3,max=255"); err != nil {
			// send error
			request(errBundle.Length(3, 255), RegisterFillStepEmpty)
			// undo
			r.FillStep--
			break
		}

		// check format
		if err := validator.V.Var(text, "fullName"); err != nil {
			// send error
			request(errBundle.FullName(), RegisterFillStepEmpty)
			// undo
			r.FillStep--
			break
		}

		// fill
		r.Manager.FullName = text
		// request next step
		request(fillBundle.Email(), RegisterFillStepEmpty)
	case RegisterFillStepEmail:
		// check email
		if err := validator.V.Var(text, "email"); err != nil {
			// send error
			request(errBundle.Email(), RegisterFillStepEmpty)
			// undo
			r.FillStep--
			break
		}
		// fill
		r.Manager.Email = text
		// request next step
		request(fillBundle.PhoneNumber(), RegisterFillStepPhoneNumber)
	case RegisterFillStepPhoneNumber:
		// skip
		if skip {
			r.Manager.PhoneNumber = nil
			// request next step
			request(fillBundle.Company(), RegisterFillStepCompany)
			break
		}

		// delete extra symbols
		text = model.FormatToE164(text)

		// check phone number
		if err := validator.V.Var(text, "e164"); err != nil {
			// send error
			request(errBundle.PhoneNumber(), RegisterFillStepEmpty)
			// undo
			r.FillStep--
			break
		}
		// fill
		r.Manager.PhoneNumber = &text
		// request next step
		request(fillBundle.Company(), RegisterFillStepCompany)
	case RegisterFillStepCompany:
		// skip
		if skip {
			r.Manager.Company = nil
			// jump to ready
			r.FillStep++
			break
		}

		// check length
		if err := validator.V.Var(text, "min=3,max=255"); err != nil {
			// send error
			request(errBundle.Length(3, 255), RegisterFillStepEmpty)
			// undo
			r.FillStep--
			break
		}
		// fill
		r.Manager.Company = &text
		// jump to ready
		r.FillStep++
	}

	return r.done(), nil
}

// Ready completes all data and send request
func (r *Register) Ready(
	managerClient manager.Client,
	send func(text string),
	bundle lang.Messages,
) error {
	// check if state done
	if !r.done() {
		return session.ErrStateNotDone
	}

	// request add parcel
	m := r.Manager
	err := managerClient.Register(context.Background(), &managerpb.RegisterRequest{
		ManagerTelegramId:  int64(m.TelegramID),
		ManagerFullName:    m.FullName,
		ManagerEmail:       m.Email,
		ManagerPhoneNumber: m.PhoneNumber,
		ManagerCompany:     m.Company,
	})
	if errors.Is(err, manager.ErrManagerIsAlreadyExist) {
		send(bundle.Common().Errors().AlreadyRegistered())
		return nil
	}
	if err != nil {
		send(bundle.Common().Errors().Internal())
		return err
	}

	// optional
	phoneNumber := ""
	if r.Manager.PhoneNumber != nil {
		phoneNumber = *r.Manager.PhoneNumber
	}

	company := ""
	if r.Manager.Company != nil {
		company = *r.Manager.Company
	}

	send(bundle.Register().Points().Ready(
		r.Manager.FullName,
		r.Manager.Email,
		phoneNumber,
		company,
	))
	return nil
}
