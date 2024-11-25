package state

import "github.com/subliker/track-parcel-service/internal/pkg/session"

type CheckParcel struct {
	TrackNum string
}

func SetCheckParcel(ss session.Session) {
	ss.SetState(CheckParcel{})
}

func (c *CheckParcel) Ended() bool {
	return c.TrackNum != ""
}
