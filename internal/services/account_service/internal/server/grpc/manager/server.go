package manager

import (
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

type ServerApi struct {
	pb.UnimplementedManagerServer

	repo   store.ManagerRepository
	logger logger.Logger
}

// New creates new instance of manager server api
func New(logger logger.Logger, repo store.ManagerRepository) *ServerApi {
	return &ServerApi{
		repo:   repo,
		logger: logger.WithFields("layer", "grpc server api", "server", "manager"),
	}
}
