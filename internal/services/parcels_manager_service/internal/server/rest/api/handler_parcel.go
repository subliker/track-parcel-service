package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
)

type (
	AddParcelRequest struct {
		ParcelName           string `json:"parcel_name"`
		ParcelRecipient      string `json:"parcel_recipient"`
		ParcelArrivalAddress string `json:"parcel_arrival_address"`
		ParcelForecatDate    string `json:"parcel_forecast_date"`
		ParcelDescription    string `json:"parcel_description"`
	}
	AddParcelResponse struct {
		ParcelTrackNumber string `json:"parcel_track_number"`
	}
)

// @Summary		Add Parcel
// @Description	Add Parcel adds new manager's parcel and returns unique parcel track number.
// @Tags			Parcels
// @Accept			json
// @Produce		json
// @Param			request	body		AddParcelRequest	true	"New Parcel data"
// @Success		201		{object}	AddParcelResponse	"Parcel was added"
// @Failure		400		{object}	string
// @Failure		500		{object}	string
// @Router			/parcels [post]
func (s *Server) handleAddParcel() http.HandlerFunc {
	logger := s.logger.WithFields("handler", "add parcel")
	const errMsg = "add parcel error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value(contextKeyManagerTID).(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		var req AddParcelRequest

		// parse json
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("invalid request body: %s", err), http.StatusBadRequest)
			return
		}

		// parse time
		forecastDate, err := time.Parse(time.RFC3339, req.ParcelForecatDate)
		if err != nil {
			http.Error(w, "invalid time format", http.StatusBadRequest)
			return
		}

		// add parcel
		trackNumber, err := s.store.Add(model.Parcel{
			Name:           req.ParcelName,
			ManagerID:      tID,
			Recipient:      req.ParcelRecipient,
			ArrivalAddress: req.ParcelArrivalAddress,
			ForecastDate:   forecastDate,
			Description:    req.ParcelDescription,
		})
		if err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// send response
		res := AddParcelResponse{
			ParcelTrackNumber: string(trackNumber),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

// @Summary		Delete Parcel
// @Description	Delete manager's Parcel with checkpoints.
// @Tags			Parcels
// @Accept			json
// @Produce		json
// @Param			track-number	path	string	true	"Parcel's track number"
// @Success		204				"Parcel was deleted"
// @Failure		400				{object}	string
// @Failure		500				{object}	string
// @Router			/parcels/{track-number} [delete]
func (s *Server) handleDeleteParcel() http.HandlerFunc {
	logger := s.logger.WithFields("handler", " delete parcel")
	const errMsg = "delete parcel error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value(contextKeyManagerTID).(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// getting track number
		trackNumber, ok := mux.Vars(r)["track-number"]
		if trackNumber == "" || !ok {
			http.Error(w, "parcel track number is not set", http.StatusBadRequest)
		}

		// check access
		ok, err := s.store.CheckAccess(model.TrackNumber(trackNumber), tID)
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

		// delete parcel
		if err := s.store.Delete(model.TrackNumber(trackNumber)); err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

type GetInfoResponse struct {
	ParcelName           string `json:"parcel_name"`
	ParcelRecipient      string `json:"parcel_recipient"`
	ParcelArrivalAddress string `json:"parcel_arrival_address"`
	ParcelForecatDate    string `json:"parcel_forecast_date"`
	ParcelDescription    string `json:"parcel_description"`
	ParcelStatus         string `json:"parcel_status"`
}

// @Summary		Get Parcel
// @Description	Get Parcel returns parcels info by track number.
// @Tags			Parcels
// @Accept			json
// @Produce		json
// @Param			track-number	path		string	true	"Parcel's track number"
// @Success		200				{object}	GetInfoResponse
// @Failure		400				{object}	string
// @Failure		403				{object}	string
// @Failure		500				{object}	string	"internal error"
// @Router			/parcels/{track-number} [get]
func (s *Server) handleGetInfo() http.HandlerFunc {
	logger := s.logger.WithFields("handler", "get info parcel")
	const errMsg = "get info parcel error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value(contextKeyManagerTID).(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// getting track number
		trackNumber, ok := mux.Vars(r)["track-number"]
		if trackNumber == "" || !ok {
			http.Error(w, "parcel track number is not set", http.StatusBadRequest)
		}

		// check access
		ok, err := s.store.CheckAccess(model.TrackNumber(trackNumber), tID)
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

		// get parcel info
		p, err := s.store.GetInfo(model.TrackNumber(trackNumber))
		if err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// send response
		res := GetInfoResponse{
			ParcelName:           p.Name,
			ParcelRecipient:      p.Recipient,
			ParcelArrivalAddress: p.ArrivalAddress,
			ParcelForecatDate:    p.ForecastDate.Format(time.RFC3339),
			ParcelDescription:    p.Description,
			ParcelStatus:         string(p.Status),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
