package main

import (
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/config"
)

func main() {
	cfg := config.Get()

	_ = app.New(cfg)
}
