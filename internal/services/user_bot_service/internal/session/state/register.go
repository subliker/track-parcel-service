package state

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/userpb"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
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
		fmt.Print(notSpecify)
		return false, ErrIncorrectNotSpecify
	}

	// increment step
	r.FillStep++

	skip := notSpecify > RegisterFillStepEmpty

	// lang bundle
	fillBundle := bundle.Register().Points()
	switch r.FillStep {
	case RegisterFillStepFullName:
		r.User.FullName = text
		request(fillBundle.Email(), 0)
	case RegisterFillStepEmail:
		r.User.Email = text
		request(fillBundle.PhoneNumber(), RegisterFillStepPhoneNumber)
	case RegisterFillStepPhoneNumber:
		if skip {
			r.User.PhoneNumber = nil
		} else {
			r.User.PhoneNumber = &text
		}

		r.FillStep++
	}

	return r.done(), nil
}

var (
	ErrUserIsAlreadyExist = errors.New("request error: user with this id is already exists")
)

// Ready completes all data and send request
func (r *Register) Ready(
	userClient user.Client,
	sendRegister func(text string),
	bundle lang.Messages,
) error {
	// check if state done
	if !r.done() {
		return ErrStateNotDone
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
		return ErrUserIsAlreadyExist
	}
	if err != nil {
		return err
	}

	// optional
	var phoneNumber string
	if u.PhoneNumber != nil {
		phoneNumber = *u.PhoneNumber
	}

	sendRegister(bundle.Register().Points().Ready(u.FullName, u.Email, phoneNumber))
	return nil
}
