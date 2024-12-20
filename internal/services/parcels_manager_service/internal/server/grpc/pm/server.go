package pm

import (
	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pmpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
)

type ServerApi struct {
	pb.UnimplementedParcelsManagerServer

	store parcel.ManagerStore

	eventProducer event.Producer

	logger logger.Logger
}

// New creates new instance of server api
func New(logger logger.Logger, store parcel.ManagerStore, eventProducer event.Producer) *ServerApi {
	return &ServerApi{
		store:         store,
		logger:        logger.WithFields("layer", "grpc server api", "server", "parcels manager"),
		eventProducer: eventProducer,
	}
}
