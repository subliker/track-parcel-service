package api

func (s *Server) initRoutes() {
	// auth middleware
	s.router.Use(s.authApiTokenMiddleware())

	// add parcel handler
	s.router.Handle("/add-parcel", s.handleAddParcel()).Methods("POST")

	s.logger.Info("routes was initialized")
}
