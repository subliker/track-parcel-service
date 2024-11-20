package app

import (
	"git.cyberzone.dev/project-tinpers/shared/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"google.golang.org/grpc"
)

type App interface {
	Run() error
}

type app struct {
	parcelServer *grpc.Server
	grpcConfig   config.GRPCConfig

	store store.Store

	logger logger.Logger
}
