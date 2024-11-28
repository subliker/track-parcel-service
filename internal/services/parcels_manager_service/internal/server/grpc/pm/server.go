package pm

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pmpb"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
)

type ServerApi struct {
	pb.UnimplementedParcelsManagerServer

	store parcel.ManagerStore

	logger logger.Logger
}

// New creates new instance of server api
func New(logger logger.Logger, store parcel.ManagerStore) *ServerApi {
	return &ServerApi{
		store:  store,
		logger: logger.WithFields("layer", "grpc server api", "server", "parcels manager"),
	}
}
