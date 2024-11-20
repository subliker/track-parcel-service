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
	logger := s.logger.WithFields("handler", "get parcel")
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
		return nil, status.Error(codes.NotFound, errMsg)
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
