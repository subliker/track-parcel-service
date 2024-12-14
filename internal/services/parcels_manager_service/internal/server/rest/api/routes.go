package api

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) initRoutes() {
	// Swagger UI
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// parcel handlers
	parcelRouter := s.router.PathPrefix("/parcels").Subrouter()
	// auth middleware
	parcelRouter.Use(s.authApiTokenMiddleware())
	parcelRouter.Handle("/", s.handleAddParcel()).Methods("POST")
	parcelRouter.Handle("/{track-number}", s.handleGetInfo()).Methods("GET")
	parcelRouter.Handle("/{track-number}", s.handleDeleteParcel()).Methods("DELETE")

	// checkpoint handlers
	checkpointRouter := s.router.PathPrefix("/checkpoints").Subrouter()
	// auth middleware
	checkpointRouter.Use(s.authApiTokenMiddleware())
	checkpointRouter.Handle("/", s.handleAddCheckpoint()).Methods("POST")

	s.logger.Info("routes was initialized")
}
