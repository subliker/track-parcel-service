module github.com/subliker/track-parcel-service/internal/services/manager_bot_service

replace github.com/subliker/track-parcel-service/internal/pkg => ../../pkg

go 1.22.2

require (
	github.com/subliker/track-parcel-service/internal/pkg v0.0.0-00010101000000-000000000000
	github.com/urfave/cli v1.22.15
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)
