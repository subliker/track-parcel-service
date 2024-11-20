package model

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
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

type Parcel struct {
	Name string

	ManagerID TelegramID

	Recipient      string
	ArrivalAddress string
	ForecastDate   time.Time

	Description string
	Status      Status
}

type Status string

const (
	StatusUnknown   Status = "unknown"
	StatusPending   Status = "pending"
	StatusInTransit Status = "in_transit"
	StatusDelivered Status = "delivered"
	StatusCanceled  Status = "canceled"
)

type Checkpoint struct {
	Time        time.Time
	Place       string
	Description string
}
