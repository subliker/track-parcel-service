package grpc

import (
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/userpb"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/server/grpc/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/server/grpc/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc"
)

func NewServer(logger logger.Logger, store store.Store) *grpc.Server {
	// making new grpc server
	s := grpc.NewServer()

	// register servers
	managerpb.RegisterManagerServer(s, manager.New(logger, store.Manager()))
	userpb.RegisterUserServer(s, user.New(logger, store.User()))
	return s
}
