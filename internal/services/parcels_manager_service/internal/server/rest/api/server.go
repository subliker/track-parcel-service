package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	docs "github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/docs"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/broker/rabbitmq/event"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
)

//	@title			Parcels Manager API
//	@version		1.0
//	@description	This server is useful for automated delivery info collectors to update parcel data in tracking system.

//	@contact.name	Shcherbachev Andrew
//	@contact.url	http://t.me/subliker
//	@contact.email	subliker0@gmail.com

//	@host		localhost:8080
//	@BasePath	/api/v1
//	@schemes	http

// @securityDefinitions.apikey	ManagerApiKey
// @in							header
// @name						Authorization
type Server struct {
	server *http.Server
	router *mux.Router

	managerClient manager.Client

	store parcel.ManagerStore

	eventProducer event.Producer

	logger logger.Logger
}

// New creates new instance of rest api server
func New(logger logger.Logger, cfg Config, managerClient manager.Client, store parcel.ManagerStore, eventProducer event.Producer) *Server {
	addr := fmt.Sprintf("localhost:%d", cfg.Port)
	s := Server{
		server: &http.Server{
			Addr: addr,
		},
		router:        mux.NewRouter().PathPrefix("/api/v1").Subrouter(),
		store:         store,
		managerClient: managerClient,
		eventProducer: eventProducer,
		logger:        logger.WithFields("layer", "rest"),
	}
	s.server.Handler = s.router

	s.initRoutes()

	// swagger options
	docs.SwaggerInfo.Host = addr

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
