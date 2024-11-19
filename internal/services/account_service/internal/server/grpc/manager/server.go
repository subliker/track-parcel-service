package manager

import (
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/manager"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

type ServerApi struct {
	pb.UnimplementedManagerServer

	store store.Store
}

// New creates new instance of manager server api
func New(store store.Store) *ServerApi {
	return &ServerApi{
		store: store,
	}
}

// func (s *ServerApi) Register(context.Context, *pb.RegisterRequest) (*emptypb.Empty, error)

// func (s *ServerApi) GetInfo(context.Context, *pb.GetInfoRequest) (*pb.GetInfoResponse, error)
