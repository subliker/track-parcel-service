package parcel

import "time"

type Parcel struct {
	ID   string
	Name string

	ManagerID string

	Recipient      string
	ArrivalAddress string
	ForecastDate   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

	Description string
	Status      Status
	Checkpoints []Checkpoint
}

type Status uint

const (
	StatusNone Status = iota
	StatusCreated
	StatusOnTheWay
	StatusDelivered
)

type Checkpoint struct {
	Time        time.Time
	Place       string
	Description string
}
