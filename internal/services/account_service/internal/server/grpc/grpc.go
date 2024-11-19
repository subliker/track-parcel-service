package grpc

import (
	managerpb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/manager"
	userpb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/server/grpc/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/server/grpc/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc"
)

func NewServer(store store.Store) *grpc.Server {
	// making new grpc server
	s := grpc.NewServer()

	// register servers
	managerpb.RegisterManagerServer(s, manager.New(store))
	userpb.RegisterUserServer(s, user.New(store))
	return s
}
