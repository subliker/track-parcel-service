package app

import (
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/config"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"google.golang.org/grpc"
)

type App interface {
	Run() error
}

type app struct {
	parcelServer *grpc.Server
	grpcConfig   config.GRPCConfig

	store parcel.Store

	logger logger.Loggers
}
