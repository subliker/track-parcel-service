package main

import (
	"flag"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
)

func main() {
	flag.Parse()

	// creating logger
	logger := zap.NewLogger()

	// reading config
	cfg :=
}
