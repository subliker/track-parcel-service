package grpc

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pmpb"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/grpc/pm"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
	"google.golang.org/grpc"
)

func NewServer(logger logger.Logger, store parcel.Store) *grpc.Server {
	// making new grpc server
	s := grpc.NewServer()

	// register server
	pmpb.RegisterParcelsManagerServer(s, pm.New(logger, store))
	return s
}
