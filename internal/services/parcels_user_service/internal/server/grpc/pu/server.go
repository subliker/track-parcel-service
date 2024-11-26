package pu

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pupb"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/store/parcel"
)

type ServerApi struct {
	pb.UnimplementedParcelsManagerServer

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
