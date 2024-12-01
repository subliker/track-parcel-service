package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/subliker/track-parcel-service/internal/pkg/model"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel"
)

func (s *Server) handleAddParcel() http.HandlerFunc {
	type Request struct {
		ParcelName           string `json:"parcel_name"`
		ParcelRecipient      string `json:"parcel_recipient"`
		ParcelArrivalAddress string `json:"parcel_arrival_address"`
		ParcelForecatDate    string `json:"parcel_forecast_date"`
		ParcelDescription    string `json:"parcel_description"`
		ParcelStatus         string `json:"parcel_status"`
	}
	type Response struct {
		ParcelTrackNumber string `json:"parcel_track_number"`
	}
	logger := s.logger.WithFields("handler", "add parcel")
	const errMsg = "add parcel error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value("manager_telegram_id").(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		var req Request

		// parse json
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// parse time
		forecastDate, err := time.Parse(time.RFC3339, req.ParcelForecatDate)
		if err != nil {
			http.Error(w, "invalid time format", http.StatusBadRequest)
			return
		}

		// parse status
		parcelStatus, ok := model.StatusValue[req.ParcelStatus]
		if !ok {
			http.Error(w, "invalid parcel status enum", http.StatusBadRequest)
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
			Status:         parcelStatus,
		})
		if err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// send response
		res := Response{
			ParcelTrackNumber: string(trackNumber),
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

func (s *Server) handleDeleteParcel() http.HandlerFunc {
	type Request struct {
		ParcelTrackNumber string `json:"parcel_track_number"`
	}
	logger := s.logger.WithFields("handler", " delete parcel")
	const errMsg = "delete parcel error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value("manager_telegram_id").(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		var req Request

		// parse json
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
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
			http.Error(w, "you havn't access for this track number", http.StatusForbidden)
			return
		}

		// delete parcel
		if err := s.store.Delete(model.TrackNumber(req.ParcelTrackNumber)); err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *Server) handleGetInfo() http.HandlerFunc {
	type Request struct {
		ParcelTrackNumber string `json:"parcel_track_number"`
	}
	type Response struct {
		ParcelName           string `json:"parcel_name"`
		ParcelRecipient      string `json:"parcel_recipient"`
		ParcelArrivalAddress string `json:"parcel_arrival_address"`
		ParcelForecatDate    string `json:"parcel_forecast_date"`
		ParcelDescription    string `json:"parcel_description"`
		ParcelStatus         string `json:"parcel_status"`
	}
	logger := s.logger.WithFields("handler", " get info parcel")
	const errMsg = "get info parcel error: %s"
	return func(w http.ResponseWriter, r *http.Request) {
		// getting managerTelegramID from middleware
		tID, ok := r.Context().Value("manager_telegram_id").(model.TelegramID)
		if !ok {
			logger.Errorf(errMsg, "middleware context value fail")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		var req Request

		// parse json
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
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
			http.Error(w, "you havn't access for this track number", http.StatusForbidden)
			return
		}

		// get parcel info
		p, err := s.store.GetInfo(model.TrackNumber(req.ParcelTrackNumber))
		if err != nil {
			logger.Errorf(errMsg, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// send response
		res := Response{
			ParcelName:           p.Name,
			ParcelRecipient:      p.Recipient,
			ParcelArrivalAddress: p.ArrivalAddress,
			ParcelForecatDate:    p.ForecastDate.Format(time.RFC3339),
			ParcelDescription:    p.Description,
			ParcelStatus:         string(p.Status),
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}
