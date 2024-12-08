package model

import (
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
)

type TrackNumber string

func NewTrackNumber() TrackNumber {
	randomString := func(alphabet string, length int) string {
		result := make([]byte, length)
		for i := 0; i < length; i++ {
			index, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
			if err != nil {
				zap.Logger.Fatalf("error get rand int in track number: %s", err)
			}
			result[i] = alphabet[index.Int64()]
		}
		return string(result)
	}

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	return TrackNumber(randomString(letters, 2) + randomString(digits, 9))
}

const ForecastDateLayout = "02.01.2006 15:04"

type Parcel struct {
	Name string `validate:"required,min=3,max=100"`

	ManagerID TelegramID `validate:"required"`

	Recipient      string    `validate:"required,min=3,max=255"`
	ArrivalAddress string    `validate:"required,min=3,max=255"`
	ForecastDate   time.Time `validate:"required"`

	Description string `validate:"required,max=255"`
	Status      Status `validate:"required"`
}

func (p *Parcel) Validate() error {
	if err := validator.V.Struct(p); err != nil {
		return err
	}

	if _, ok := StatusValue[string(p.Status)]; !ok {
		return errors.New("status enum value doesn't exist")
	}
	return nil
}

type Status string

const (
	StatusUnknown   Status = "UNKNOWN"
	StatusPending   Status = "PENDING"
	StatusInTransit Status = "IN_TRANSIT"
	StatusDelivered Status = "DELIVERED"
	StatusCanceled  Status = "CANCELED"
)

var StatusValue = map[string]Status{
	"UNKNOWN":    StatusUnknown,
	"PENDING":    StatusPending,
	"IN_TRANSIT": StatusInTransit,
	"DELIVERED":  StatusDelivered,
	"CANCELED":   StatusCanceled,
}

type Checkpoint struct {
	Time         time.Time
	Place        string
	Description  string
	ParcelStatus Status
}
