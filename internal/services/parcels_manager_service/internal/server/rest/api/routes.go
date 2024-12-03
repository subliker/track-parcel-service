package api

func (s *Server) initRoutes() {
	// auth middleware
	s.router.Use(s.authApiTokenMiddleware())

	// add parcel handler
	s.router.Handle("/parcels", s.handleAddParcel()).Methods("POST")
	s.router.Handle("/parcels/{track-number}", s.handleGetInfo()).Methods("GET")
	s.router.Handle("/parcels/{track-number}", s.handleDeleteParcel()).Methods("DELETE")

	s.logger.Info("routes was initialized")
}
