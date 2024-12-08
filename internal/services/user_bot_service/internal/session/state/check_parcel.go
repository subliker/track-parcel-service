package state

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pu"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/session"
	"github.com/subliker/track-parcel-service/internal/services/user_bot_service/internal/lang"
)

// CheckParcel is state to request parcel from parcels user service
type CheckParcel struct {
	TrackNum model.TrackNumber
}

// SetCheckParcel sets empty check parcel state in user session
func SetCheckParcel(ss session.Session) {
	ss.SetState(CheckParcel{})
}

// done returns true if all state data is filled up
func (c *CheckParcel) done() bool {
	return c.TrackNum != ""
}

// Next starts one of iterations to fill state up and returns true if all state data is filled up
func (c *CheckParcel) Next(text string) (bool, error) {
	c.TrackNum = model.TrackNumber(text)
	return c.done(), nil
}

// Ready completes all data and send request
func (c *CheckParcel) Ready(
	parcelsUserClient pu.Client,
	sendParcel func(text string, subscribed bool),
	bundle lang.Messages,
	userTID model.TelegramID,
) error {
	// check if state done
	if !c.done() {
		return session.ErrStateNotDone
	}

	// request parcel
	res, err := parcelsUserClient.GetParcel(context.Background(), &pupb.GetParcelRequest{
		TrackNumber:    string(c.TrackNum),
		UserTelegramId: int64(userTID),
	})
	if errors.Is(err, pu.ErrParcelNotFound) {
		return session.ErrResNotFound
	}
	if err != nil {
		return err
	}

	// enum convert
	parcelStatus, ok := model.StatusValue[res.Parcel.Status.String()]
	if !ok {
		return fmt.Errorf("parcel response status incorrect value: %s", res.Parcel.Status.String())
	}

	// show parcel
	sendParcel(bundle.CheckParcel().Main(
		res.Parcel.Name,
		res.Parcel.Recipient,
		res.Parcel.ArrivalAddress,
		res.Parcel.ForecastDate.AsTime().Format(model.ForecastDateLayout),
		res.Parcel.Description,
		string(parcelStatus),
	)+"\n"+
		bundle.CheckParcel().Subscription(res.UserSubscribed), res.UserSubscribed)
	return nil
}
