package pm

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/subliker/track-parcel-service/internal/pkg/gen/pmpb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerApi) AddCheckpoint(ctx context.Context, req *pb.AddCheckpointRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "add checkpoint")
	const errMsg = "error add checkpoint for parcel(%s): %s"

	// add checkpoint to store
	err := s.store.AddCheckpoint(model.TrackNumber(req.TrackNumber), model.Checkpoint{
		Time:        req.Time.AsTime(),
		Place:       req.Place,
		Description: req.Description,
	})
	if errors.Is(err, parcel.ErrIncorrectForeignTrackNumber) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) GetCheckpoints(ctx context.Context, req *pb.GetCheckpointsRequest) (*pb.GetCheckpointsResponse, error) {
	logger := s.logger.WithFields("handler", "get checkpoints")
	const errMsg = "error get checkpoints for parcel(%s): %s"

	// check if parcel exists
	exists, err := s.store.Exists(model.TrackNumber(req.TrackNumber))
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}
	if !exists {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, parcel.ErrParcelNotFound)
		return nil, status.Error(codes.NotFound, errMsg)
	}

	// get checkpoints from store
	cps, err := s.store.GetCheckpoints(model.TrackNumber(req.TrackNumber), req.Page, req.PageSize)
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	// transfer to proto checkpoints
	protoCps := make([]*pb.Checkpoint, len(cps))
	for i := range protoCps {
		cp := cps[i]
		protoCps[i] = &pb.Checkpoint{
			Time:        timestamppb.New(cp.Time),
			Place:       cp.Place,
			Description: cp.Description,
		}
	}

	return &pb.GetCheckpointsResponse{
		Checkpoints: protoCps,
	}, nil
}
