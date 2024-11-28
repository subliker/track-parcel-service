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
	config        config.RESTConfig
	router        *mux.Router
	managerClient manager.Client
	store         parcel.ManagerStore

	logger logger.Logger
}

// New creates new instance of rest api server
func New(logger logger.Logger, cfg config.RESTConfig, managerClient manager.Client, store parcel.ManagerStore) *Server {
	s := Server{
		config:        cfg,
		router:        mux.NewRouter(),
		store:         store,
		managerClient: managerClient,
		logger:        logger.WithFields("layer", "rest"),
	}

	s.logger.Info("rest api server instance created")
	return &s
}

// Run runs server listening
func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router)
}
