package pu

import (
	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/store/parcel"
)

type ServerApi struct {
	pb.UnimplementedParcelsUserServer

	store  parcel.UserStore
	logger logger.Logger
}

// New creates new instance of server api
func New(logger logger.Logger, store parcel.UserStore) *ServerApi {
	return &ServerApi{
		store:  store,
		logger: logger.WithFields("layer", "grpc server api", "server", "parcels user"),
	}
}
