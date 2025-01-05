package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/gen/notificationpb"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/parcelpb"
	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	AddCheckpointRequest struct {
		ParcelTrackNumber      string `json:"parcel_track_number"`
		CheckpointTime         string `json:"checkpoint_time"`
		CheckpointPlace        string `json:"checkpoint_place"`
		CheckpointDescription  string `json:"checkpoint_description"`
		CheckpointParcelStatus string `json:"checkpoint_parcel_status"`
	}
)

// @Summary Add Checkpoint
// @Description Add Checkpoint adds new parcel's and send event for notification system
// @Tags Parcels Checkpoints
// @Accept json
// @Produce json
// @Param track-number path string true
// @Success 201 "checkpoint was added"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /checkpoints [post]
func (s *Server) handleAddCheckpoint() http.HandlerFunc {
	logger := s.logger.WithFields("handler", "add checkpoint")
	const errMsg = "add checkpoint error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value(contextKeyManagerTID).(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		var req AddCheckpointRequest

		// parse json
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("invalid request body: %s", err), http.StatusBadRequest)
			return
		}

		// check access
		ok, err := s.store.CheckAccess(model.TrackNumber(req.ParcelTrackNumber), tID)
		if err == parcel.ErrParcelNotFound {
			http.Error(w, "parcel with this track number is not found", http.StatusNotFound)
		}
		if err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Error(w, "you haven't access for this track number", http.StatusForbidden)
			return
		}

		// parse time
		cTime, err := time.Parse(time.RFC3339, req.CheckpointTime)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid time format: requires %s", time.RFC3339), http.StatusBadRequest)
			return
		}

		// parse status
		parcelStatus, ok := model.StatusValue[req.CheckpointParcelStatus]
		if !ok {
			http.Error(w, "invalid parcel status enum", http.StatusBadRequest)
			return
		}

		// add checkpoint
		err = s.store.AddCheckpoint(model.TrackNumber(req.ParcelTrackNumber), model.Checkpoint{
			Time:         cTime,
			Place:        req.CheckpointPlace,
			Description:  req.CheckpointDescription,
			ParcelStatus: parcelStatus,
		})
		if err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// parse status for proto
		pStatus, ok := parcelpb.Status_value[req.ParcelTrackNumber]
		if !ok {
			s.logger.Error("error parsing proto parcel status, enum does not match in proto and model")

			w.WriteHeader(http.StatusCreated)
		}

		// publishing event
		if err := s.eventProducer.Publish(&notificationpb.Event{
			TrackNumber: req.ParcelTrackNumber,
			Checkpoint: &parcelpb.Checkpoint{
				Time:         timestamppb.New(cTime),
				Place:        req.CheckpointPlace,
				Description:  req.CheckpointDescription,
				ParcelStatus: parcelpb.Status(pStatus),
			},
		}); err != nil {
			logger.Errorf("error publishing checkpoint event: %s", err)
		}

		w.WriteHeader(http.StatusCreated)
	}
}
