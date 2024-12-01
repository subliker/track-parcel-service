package grpc

import (
	"github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/server/grpc/pu"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/store/parcel"
	"google.golang.org/grpc"
)

func NewServer(logger logger.Logger, store parcel.UserStore) *grpc.Server {
	// making new grpc server
	s := grpc.NewServer()

	// register server
	pupb.RegisterParcelsUserServer(s, pu.New(logger, store))
	return s
}
