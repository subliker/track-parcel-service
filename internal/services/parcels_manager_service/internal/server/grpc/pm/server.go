package pm

import (
	"context"
	"errors"
	"fmt"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	pb "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/pm"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerApi struct {
	pb.UnimplementedParcelsManagerServer

	store  parcel.Store
	logger logger.Logger
}

// New creates new instance of server api
func New(logger logger.Logger, store parcel.Store) *ServerApi {
	return &ServerApi{
		store:  store,
		logger: logger.WithFields("layer", "grpc server api"),
	}
}

func (s *ServerApi) AddParcel(ctx context.Context, req *pb.AddParcelRequest) (*pb.AddParcelResponse, error) {
	logger := s.logger.WithFields("handler", "add parcel")
	const errMsg = "error add parcel: %s"

	// add parcel to store
	trackNum, err := s.store.Add(model.Parcel{
		Name:           req.ParcelName,
		ManagerID:      model.TelegramID(req.ManagerTelegramId),
		Recipient:      req.ParcelRecipient,
		ArrivalAddress: req.ParcelArrivalAddress,
		ForecastDate:   req.ParcelForecastDate.AsTime(),
		Description:    req.ParcelDescription,
		Status:         model.Status(req.ParcelStatus),
	})
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, err)
		logger.Error(errMsg)
		return &pb.AddParcelResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &pb.AddParcelResponse{
		TrackNumber: string(trackNum),
	}, nil
}

func (s *ServerApi) DeleteParcel(ctx context.Context, req *pb.DeleteParcelRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "delete parcel")
	const errMsg = "error delete parcel(%s): %s"

	// delete parcel from store
	err := s.store.Delete(model.TrackNumber(req.TrackNumber))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) GetParcelInfo(ctx context.Context, req *pb.GetParcelRequest) (*pb.GetParcelResponse, error) {
	logger := s.logger.WithFields("handler", "get parcel info")
	const errMsg = "error get parcel(%s): %s"

	// get parcel from store
	p, err := s.store.GetInfo(model.TrackNumber(req.TrackNumber))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.GetParcelResponse{
		ParcelName:           p.Name,
		ManagerTelegramId:    int64(p.ManagerID),
		ParcelRecipient:      p.Recipient,
		ParcelArrivalAddress: p.ArrivalAddress,
		ParcelForecastDate:   timestamppb.New(p.ForecastDate),
		ParcelDescription:    p.Description,
		ParcelStatus:         pb.ParcelStatus(pb.ParcelStatus_value[string(p.Status)]),
	}, nil
}

func (s *ServerApi) AddCheckpoint(ctx context.Context, req *pb.AddCheckpointRequest) (*emptypb.Empty, error) {
	logger := s.logger.WithFields("handler", "add checkpoint")
	const errMsg = "error add checkpoint for parcel(%d): %s"

	// add checkpoint to store
	err := s.store.AddCheckpoint(model.TrackNumber(req.TrackNumber), model.Checkpoint{
		Time:        req.Time.AsTime(),
		Place:       req.Place,
		Description: req.Description,
	})
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
	}

	return nil, nil
}

func (s *ServerApi) GetCheckpoints(ctx context.Context, req *pb.GetCheckpointsRequest) (*pb.GetCheckpointsResponse, error) {
	logger := s.logger.WithFields("handler", "get checkpoints")
	const errMsg = "error get checkpoints for parcel(%s): %s"

	// get checkpoints from store
	cps, err := s.store.GetCheckpoints(model.TrackNumber(req.TrackNumber), int(req.Page), int(req.PageSize))
	if errors.Is(err, parcel.ErrParcelNotFound) {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	if err != nil {
		errMsg := fmt.Sprintf(errMsg, req.TrackNumber, err)
		logger.Error(errMsg)
		return nil, status.Error(codes.NotFound, errMsg)
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
