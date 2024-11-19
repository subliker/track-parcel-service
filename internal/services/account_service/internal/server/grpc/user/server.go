package user

import (
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/account/user"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
)

type ServerApi struct {
	pb.UnimplementedUserServer

	store store.Store
}

// New creates new instance of user server api
func New(store store.Store) *ServerApi {
	return &ServerApi{
		store: store,
	}
}

// func (s *ServerApi) Register(context.Context, *pb.RegisterRequest) (*emptypb.Empty, error) {

// }

// func (s *ServerApi) GetInfo(context.Context, *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {

// }
