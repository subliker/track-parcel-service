package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
)

type Server struct {
	server *http.Server
	router *mux.Router

	managerClient manager.Client

	store parcel.ManagerStore

	logger logger.Logger
}

// New creates new instance of rest api server
func New(logger logger.Logger, cfg config.RESTConfig, managerClient manager.Client, store parcel.ManagerStore) *Server {
	s := Server{
		server: &http.Server{
			Addr: fmt.Sprintf("localhost:%d", cfg.Port),
		},
		router:        mux.NewRouter(),
		store:         store,
		managerClient: managerClient,
		logger:        logger.WithFields("layer", "rest"),
	}
	s.server.Handler = s.router

	s.initRoutes()

	s.logger.Info("rest api server instance created")
	return &s
}

// Run runs server listening
func (s *Server) Run() error {
	s.logger.Infof("server starting on address: %s", s.server.Addr)
	return s.server.ListenAndServe()
}

// Close closes rest api server
func (s *Server) Close() error {
	if err := s.server.Close(); err != nil {
		return err
	}

	s.logger.Info("server closed")
	return nil
}
