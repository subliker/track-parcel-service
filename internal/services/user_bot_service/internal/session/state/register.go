package state

import (
	"context"
	"errors"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/userpb"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/lang"
)

// Register is state to request register user from user service
type Register struct {
	User     model.User
	FillStep RegisterFillStep
}

// RegisterFillStep is enum to mention every fill step
type RegisterFillStep uint

const (
	RegisterFillStepEmpty RegisterFillStep = iota
	RegisterFillStepFullName
	RegisterFillStepEmail
	RegisterFillStepPhoneNumber
	RegisterFillStepReady
)

// SetRegister sets empty register state in user session
func SetRegister(ss session.Session, tID model.TelegramID) {
	ss.SetState(Register{
		User: model.User{
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
		// fill
		r.User.FullName = text
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
		r.User.Email = text
		// request next step
		request(fillBundle.PhoneNumber(), RegisterFillStepPhoneNumber)
	case RegisterFillStepPhoneNumber:
		// skip
		if skip {
			r.User.PhoneNumber = nil
			// jump to ready
			r.FillStep++
			break
		}

		// check phone number
		if err := validator.V.Var(text, "e164"); err != nil {
			// send error
			request(errBundle.PhoneNumber(), RegisterFillStepEmpty)
			// undo
			r.FillStep--
			break
		}
		// fill
		r.User.PhoneNumber = &text
		// jump to ready
		r.FillStep++
	}

	return r.done(), nil
}

// Ready completes all data and send request
func (r *Register) Ready(
	userClient user.Client,
	send func(text string),
	bundle lang.Messages,
) error {
	// check if state done
	if !r.done() {
		return session.ErrStateNotDone
	}

	// request register user
	u := r.User
	err := userClient.Register(context.Background(), &userpb.RegisterRequest{
		UserTelegramId:  int64(u.TelegramID),
		UserFullName:    u.FullName,
		UserEmail:       u.Email,
		UserPhoneNumber: u.PhoneNumber,
	})
	if errors.Is(err, user.ErrUserIsAlreadyExist) {
		return errors.New("request error: user with this id is already exists")
	}
	if err != nil {
		return err
	}

	// optional
	phoneNumber := ""
	if u.PhoneNumber != nil {
		phoneNumber = *u.PhoneNumber
	}

	send(bundle.Register().Points().Ready(u.FullName, u.Email, phoneNumber))
	return nil
}
