package main

import (
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/app"
	"github.com/subliker/track-parcel-service/internal/services/manager_bot_service/internal/config"
)

func main() {
	// reading config
	cfg := config.Get()

	// creating new instance of app
	a := app.New(cfg)
	// running app
	a.Run()
}
