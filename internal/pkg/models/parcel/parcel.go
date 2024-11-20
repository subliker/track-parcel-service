package parcel

import "time"

type Parcel struct {
	ID   string
	Name string

	ManagerID string

	Recipient      string
	ArrivalAddress string
	ForecastDate   time.Time

	Description string
	Status      Status
	Checkpoints []Checkpoint
}

type Status uint

const (
	StatusUnknown Status = iota
	StatusPending
	StatusInTransit
	StatusDelivered
	StatusCanceled
)

type Checkpoint struct {
	Time        time.Time
	Place       string
	Description string
}
