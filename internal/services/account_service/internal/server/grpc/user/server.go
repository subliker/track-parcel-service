package user

import (
	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/account/userpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

type ServerApi struct {
	pb.UnimplementedUserServer

	repo   store.UserRepository
	logger logger.Logger
}

// New creates new instance of user server api
func New(logger logger.Logger, repo store.UserRepository) *ServerApi {
	logger = logger.WithFields("layer", "grpc server api", "server", "user")
	return &ServerApi{
		repo:   repo,
		logger: logger,
	}
}
